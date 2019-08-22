package homebrew

import (
	"context"
	"io"
	"os"
)

const (
	rootCmd      = "brew"
	installCmd   = "install"
	uninstallCmd = "uninstall"
	listCmd      = "list"
	searchCmd    = "search"
)

// Homebrew is a client to run Homebrew command.
type Homebrew interface {
	// Install executes `brew install [options] formula`.
	Install(ctx context.Context, formula string, opts ...InstallOption) error
	// Uninstall executes `brew uninstall, rm, remove [options] formula`.
	Uninstall(ctx context.Context, formula string, opts ...UninstallOption) error
	// List executes `brew list, ls [options] [formula]`.
	List(ctx context.Context, opts ...ListOption) error
	// Search executes `brew search [options] [text|/text/]`.
	Search(ctx context.Context, text string, opts ...SearchOption) error
}

type homebrew struct {
	stdout io.Writer
	stderr io.Writer
}

var _ Homebrew = (*homebrew)(nil)

type Option func(*homebrew)

// WithStdout applied a custom io.Writer as stdout to Homebrew.
func WithStdout(stdout io.Writer) Option {
	return func(h *homebrew) {
		h.stdout = stdout
	}
}

// WithStderr applied a custom io.Writer as stderr to Homebrew.
func WithStderr(stderr io.Writer) Option {
	return func(h *homebrew) {
		h.stderr = stderr
	}
}

// New instantiates Homebrew with some options.
func New(opts ...Option) Homebrew {
	h := &homebrew{
		stdout: os.Stdout,
		stderr: os.Stderr,
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}
