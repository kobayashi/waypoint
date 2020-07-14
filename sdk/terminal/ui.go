package terminal

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"text/tabwriter"

	"github.com/fatih/color"
)

const (
	TermRows    = 10
	TermColumns = 100
)

// Passed to UI.NamedValues to provide a nicely formatted key: value output
type NamedValue struct {
	Name  string
	Value interface{}
}

// Passed to UI.Table to provide a nicely formatted table.
type Table struct {
	Headers []string
	Rows    [][]TableEntry
}

type TableHeader struct {
	Name  string
	Color string
}

func NewTable(headers ...string) *Table {
	return &Table{
		Headers: headers,
	}
}

const (
	Yellow = "yellow"
	Green  = "green"
	Red    = "red"
)

type TableEntry struct {
	Value string
	Color string
}

func (t *Table) Rich(cols []string, colors []string) {
	var row []TableEntry

	for i, col := range cols {
		if i < len(colors) {
			row = append(row, TableEntry{Value: col, Color: colors[i]})
		} else {
			row = append(row, TableEntry{Value: col})
		}
	}

	t.Rows = append(t.Rows, row)
}

// UI is the primary interface for interacting with a user via the CLI.
//
// NOTE(mitchellh): This is an interface and not a struct directly so that
// we can support other user interaction patterns in the future more easily.
// Most importantly what I'm thinking of is when we support multiple "apps"
// in a single config file, we can build a UI that locks properly and so on
// without changing the API.
type UI interface {
	// Output outputs a message directly to the terminal. The remaining
	// arguments should be interpolations for the format string. After the
	// interpolations you may add Options.
	Output(string, ...interface{})

	// Output data as a table of data. Each entry is a row which will be output
	// with the columns lined up nicely.
	NamedValues([]NamedValue, ...Option)

	// OutputWriters returns stdout and stderr writers. These are usually
	// but not always TTYs. This is useful for subprocesses, network requests,
	// etc. Note that writing to these is not thread-safe by default so
	// you must take care that there is only ever one writer.
	OutputWriters() (stdout, stderr io.Writer, err error)

	// Status returns a live-updating status that can be used for single-line
	// status updates that typically have a spinner or some similar style.
	// While a Status is live (Close isn't called), Output should NOT be called.
	Status() Status

	// Table outputs the information formatted into a Table structure.
	Table(*Table, ...Option)

	StepGroup() StepGroup
}

type StepGroup interface {
	// Start a step in the output
	Add(string, ...interface{}) Step

	Wait()
}

type Step interface {
	TermOutput() io.Writer
	Update(string, ...interface{})
	Status(status string)
	Done()
	Abort()
}

// BasicUI
type BasicUI struct {
	status *spinnerStatus
}

// Interpret decomposes the msg and arguments into the message, style, and writer
func Interpret(msg string, raw ...interface{}) (string, string, io.Writer) {
	// Build our args and options
	var args []interface{}
	var opts []Option
	for _, r := range raw {
		if opt, ok := r.(Option); ok {
			opts = append(opts, opt)
		} else {
			args = append(args, r)
		}
	}

	// Build our message
	msg = fmt.Sprintf(msg, args...)

	// Build our config and set our options
	cfg := &config{Original: msg, Message: msg, Writer: color.Output}
	for _, opt := range opts {
		opt(cfg)
	}

	return msg, cfg.Style, cfg.Writer
}

// Output implements UI
func (ui *BasicUI) Output(msg string, raw ...interface{}) {
	msg, style, w := Interpret(msg, raw...)

	switch style {
	case HeaderStyle:
		msg = colorHeader.Sprintf("==> %s", msg)
	case ErrorStyle:
		msg = colorError.Sprint(msg)
	case WarningStyle:
		msg = colorWarning.Sprint(msg)
	case SuccessStyle:
		msg = colorSuccess.Sprint(msg)
	case InfoStyle:
		lines := strings.Split(msg, "\n")
		for i, line := range lines {
			lines[i] = colorInfo.Sprintf("    %s", line)
		}

		msg = strings.Join(lines, "\n")
	}

	st := ui.status

	if st != nil {
		st.Pause()
		defer st.Start()
	}

	// Write it
	fmt.Fprintln(w, msg)
}

func (ui *BasicUI) NamedValues(rows []NamedValue, opts ...Option) {
	cfg := &config{Writer: color.Output}
	for _, opt := range opts {
		opt(cfg)
	}

	cfg.Writer.Write([]byte{'\n'})

	tr := tabwriter.NewWriter(cfg.Writer, 1, 8, 0, ' ', tabwriter.AlignRight)
	for _, row := range rows {
		colorInfo.Fprintf(tr, "%s: \t%s\n", row.Name, row.Value)
	}

	tr.Flush()

	cfg.Writer.Write([]byte{'\n'})
}

// OutputWriters implements UI
func (ui *BasicUI) OutputWriters() (io.Writer, io.Writer, error) {
	return os.Stdout, os.Stderr, nil
}

// Status implements UI
func (ui *BasicUI) Status() Status {
	if ui.status == nil {
		ui.status = newSpinnerStatus()
	}

	return ui.status
}

type fancyStep struct {
	sg  *fancyStepGroup
	ent *DisplayEntry

	done bool

	term *Term
}

func (f *fancyStep) TermOutput() io.Writer {
	if f.term == nil {
		t, err := NewTerm(f.sg.ctx, f.ent, TermRows, TermColumns)
		if err != nil {
			panic(err)
		}

		f.term = t
	}

	return f.term
}

func (f *fancyStep) Update(str string, args ...interface{}) {
	f.ent.Update(str, args...)
}

func (f *fancyStep) Status(status string) {
	f.ent.SetStatus(status)
}

func (f *fancyStep) Done() {
	if f.done {
		return
	}

	f.ent.StopSpinner()
	f.Status(StatusOK)
	f.done = true
	f.sg.wg.Done()
}

func (f *fancyStep) Abort() {
	if f.done {
		return
	}

	f.ent.StopSpinner()
	f.Status(StatusError)

	f.done = true
	f.sg.wg.Done()
}

type fancyStepGroup struct {
	ctx    context.Context
	cancel func()

	display *Display

	wg sync.WaitGroup
}

// Start a step in the output
func (f *fancyStepGroup) Add(str string, args ...interface{}) Step {
	f.wg.Add(1)

	ent := f.display.NewStatus(0)

	ent.StartSpinner()
	ent.Update(str, args...)

	return &fancyStep{
		sg:  f,
		ent: ent,
	}
}

func (f *fancyStepGroup) Wait() {
	f.wg.Wait()
	f.cancel()

	f.display.Close()
}

func (ui *BasicUI) StepGroup() StepGroup {
	ctx, cancel := context.WithCancel(context.Background())
	display := NewDisplay(ctx, color.Output)

	return &fancyStepGroup{
		ctx:     ctx,
		cancel:  cancel,
		display: display,
	}
}

const (
	HeaderStyle  = "header"
	ErrorStyle   = "error"
	WarningStyle = "warning"
	InfoStyle    = "info"
	SuccessStyle = "success"
)

type config struct {
	// Original is the original message, this should NOT be modified.
	Original string

	// Message is the message to write.
	Message string

	// Writer is where the message will be written to.
	Writer io.Writer

	// The style the output should take on
	Style string
}

// Option controls output styling.
type Option func(*config)

// WithHeaderStyle styles the output like a header denoting a new section
// of execution. This should only be used with single-line output. Multi-line
// output will not look correct.
func WithHeaderStyle() Option {
	return func(c *config) {
		c.Style = HeaderStyle
	}
}

// WithInfoStyle styles the output like it's formatted information.
func WithInfoStyle() Option {
	return func(c *config) {
		c.Style = InfoStyle
	}
}

// WithErrorStyle styles the output as an error message.
func WithErrorStyle() Option {
	return func(c *config) {
		c.Style = ErrorStyle
	}
}

// WithWarningStyle styles the output as an error message.
func WithWarningStyle() Option {
	return func(c *config) {
		c.Style = WarningStyle
	}
}

// WithSuccessStyle styles the output as a success message.
func WithSuccessStyle() Option {
	return func(c *config) {
		c.Style = SuccessStyle
	}
}

func WithStyle(style string) Option {
	return func(c *config) {
		c.Style = style
	}
}

// WithWriter specifies the writer for the output.
func WithWriter(w io.Writer) Option {
	return func(c *config) { c.Writer = w }
}

var (
	colorHeader  = color.New(color.Bold)
	colorInfo    = color.New()
	colorError   = color.New(color.FgRed)
	colorSuccess = color.New(color.FgGreen)
	colorWarning = color.New(color.FgYellow)
)
