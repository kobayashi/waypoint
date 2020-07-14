package pack

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/buildpacks/pack"
	"github.com/buildpacks/pack/logging"
	"github.com/hashicorp/waypoint/internal/assets"
	"github.com/hashicorp/waypoint/internal/pkg/epinject"
	"github.com/hashicorp/waypoint/sdk/component"
	"github.com/hashicorp/waypoint/sdk/terminal"
)

// Builder uses `pack` -- the frontend for CloudNative Buildpacks -- to build
// an artifact from source.
type Builder struct {
	config BuilderConfig
}

// BuildFunc implements component.Builder
func (b *Builder) BuildFunc() interface{} {
	return b.Build
}

// Config is the configuration structure for the registry.
type BuilderConfig struct {
	// Control whether or not to inject the entrypoint binary into the resulting image
	DisableCEB bool `hcl:"disable_ceb,optional"`

	// The Buildpack builder image to use, defaults to the standard heroku one.
	Builder string `hcl:"builder,optional"`
}

const DefaultBuilder = "heroku/buildpacks:18"

// Config implements Configurable
func (b *Builder) Config() (interface{}, error) {
	return &b.config, nil
}

// Build
func (b *Builder) Build(
	ctx context.Context,
	ui terminal.UI,
	src *component.Source,
) (*DockerImage, error) {
	builder := b.config.Builder
	if builder == "" {
		builder = DefaultBuilder
	}

	ui.Output("Creating new buildpack-based image using builder: %s", builder, terminal.WithInfoStyle())

	sg := ui.StepGroup()

	step := sg.Add("Creating pack client")
	defer step.Abort()

	build := sg.Add("Building image")
	defer build.Abort()

	log := logging.New(build.TermOutput())

	client, err := pack.NewClient(pack.WithLogger(log))
	if err != nil {
		return nil, err
	}

	step.Done()

	err = client.Build(ctx, pack.BuildOptions{
		Image:   src.App,
		Builder: builder,
		AppPath: src.Path,
	})

	if err != nil {
		return nil, err
	}

	build.Done()

	if !b.config.DisableCEB {
		inject := sg.Add("Injecting entrypoint binary to image")
		defer inject.Abort()

		tmpdir, err := ioutil.TempDir("", "waypoint")
		if err != nil {
			return nil, err
		}

		defer os.RemoveAll(tmpdir)

		err = assets.RestoreAsset(tmpdir, "ceb/ceb")
		if err != nil {
			return nil, err
		}

		err = epinject.AlterEntrypoint(ctx, src.App+":latest", func(cur []string) (*epinject.NewEntrypoint, error) {
			ep := &epinject.NewEntrypoint{
				Entrypoint: append([]string{"/bin/wpceb"}, cur...),
				InjectFiles: map[string]string{
					filepath.Join(tmpdir, "ceb/ceb"): "/bin/wpceb",
				},
			}

			return ep, nil
		})

		if err != nil {
			return nil, err
		}

		inject.Done()
	}

	sg.Wait()

	ui.Output("Generated new Docker image: %s:latest", src.App)

	// We don't even need to inspect Docker to verify we have the image.
	// If `pack` succeeded we can assume that it created an image for us.
	return &DockerImage{
		Image: src.App,
		Tag:   "latest", // It always tags latest
	}, nil
}
