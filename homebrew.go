package homebrew

import (
	"context"
	"fmt"
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

type Homebrew interface {
	Install(ctx context.Context, formula string, opts ...InstallOption) error
	Uninstall(ctx context.Context, formula string, opts ...UninstallOption) error
	List(ctx context.Context, opts ...ListOption) error
	Search(ctx context.Context, text string, opts ...SearchOption) error
}

type homebrew struct {
	stdout io.Writer
	stderr io.Writer
}

var _ Homebrew = (*homebrew)(nil)

type Option func(*homebrew)

func WithStdout(stdout io.Writer) Option {
	return func(h *homebrew) {
		h.stdout = stdout
	}
}

func WithStderr(stderr io.Writer) Option {
	return func(h *homebrew) {
		h.stderr = stderr
	}
}

func NewHomebrew(opts ...Option) Homebrew {
	h := &homebrew{
		stdout: os.Stdout,
		stderr: os.Stderr,
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

// cmdLongOption makes a long option text from the name.
// e.g.) verbose -> --verbose
func cmdLongOption(optName string) string {
	return fmt.Sprintf("--%s", optName)
}

// cmdShortOption makes a short option text from the name.
// e.g.) l -> -l
func cmdShortOption(optName string) string {
	return fmt.Sprintf("-%s", optName)
}
