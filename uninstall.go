package homebrew

import (
	"context"
	"os/exec"
)

type UninstallOption string

const (
	UninstallOptionForce              UninstallOption = "--force"
	UninstallOptionDebug              UninstallOption = "--debug"
	UninstallOptionIgnoreDependencies UninstallOption = "--ignore-dependencies"
	UninstallOptionHelp               UninstallOption = "--help"
)

func (o UninstallOption) String() string {
	return string(o)
}

func (h *homebrew) Uninstall(ctx context.Context, formula string, opts ...UninstallOption) error {
	args := make([]string, 0, len(opts)+2)
	args = append(args, uninstallCmd)
	args = append(args, uninstallOptions(opts)...)
	args = append(args, formula)

	cmd := exec.CommandContext(ctx, rootCmd, args...)
	cmd.Stdout = h.stdout
	cmd.Stderr = h.stderr

	return cmd.Run()
}

func uninstallOptions(opts []UninstallOption) []string {
	execOpts := make([]string, 0, len(opts))
	for _, opt := range opts {
		execOpts = append(execOpts, opt.String())
	}
	return execOpts
}
