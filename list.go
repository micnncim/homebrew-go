package homebrew

import (
	"context"
	"os/exec"
)

type ListOption string

const (
	ListOptionFullName        ListOption = "--full-name"
	ListOptionUnbrewed        ListOption = "--unbrewed"
	ListOptionVersions        ListOption = "--versions"
	ListOptionMultiple        ListOption = "--multiple"
	ListOptionPinned          ListOption = "--pinned"
	ListOptionVerbose         ListOption = "--verbose"
	ListOptionDebug           ListOption = "--debug"
	ListOptionHelp            ListOption = "--help"
	ListOptionOneEntryPerLine ListOption = "-1"
	ListOptionLong            ListOption = "-l"
	ListOptionReverse         ListOption = "-r"
	ListOptionSortByTime      ListOption = "-t"
)

func (o ListOption) String() string {
	return string(o)
}

func (h *homebrew) List(ctx context.Context, opts ...ListOption) error {
	args := make([]string, 0, len(opts)+1)
	args = append(args, listCmd)
	args = append(args, listOptions(opts)...)

	cmd := exec.CommandContext(ctx, rootCmd, args...)
	cmd.Stdout = h.stdout
	cmd.Stderr = h.stderr

	return cmd.Run()
}

func listOptions(opts []ListOption) []string {
	execOpts := make([]string, 0, len(opts))
	for _, opt := range opts {
		execOpts = append(execOpts, opt.String())
	}
	return execOpts
}
