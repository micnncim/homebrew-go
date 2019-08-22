package homebrew

import (
	"context"
	"os/exec"
)

type ListOption string

const (
	ListOptionFullName        ListOption = "full-name"
	ListOptionUnbrewed        ListOption = "unbrewed"
	ListOptionVersions        ListOption = "versions"
	ListOptionMultiple        ListOption = "multiple"
	ListOptionPinned          ListOption = "pinned"
	ListOptionVerbose         ListOption = "verbose"
	ListOptionDebug           ListOption = "debug"
	ListOptionHelp            ListOption = "help"
	ListOptionOneEntryPerLine ListOption = "1"
	ListOptionLong            ListOption = "l"
	ListOptionReverse         ListOption = "r"
	ListOptionSortByTime      ListOption = "t"
)

func (o ListOption) String() string {
	return string(o)
}

func (h *homebrew) List(ctx context.Context, opts ...ListOption) error {
	args := append([]string{listCmd}, listOptions(opts)...)

	cmd := exec.CommandContext(ctx, rootCmd, args...)
	cmd.Stdout = h.stdout
	cmd.Stderr = h.stderr

	return cmd.Run()
}

func listOptions(opts []ListOption) []string {
	execOpts := make([]string, 0, len(opts))
	for _, opt := range opts {
		var execOpt string
		switch opt {
		case ListOptionOneEntryPerLine:
			execOpt = cmdShortOption(opt.String())
		case ListOptionLong:
			execOpt = cmdShortOption(opt.String())
		case ListOptionReverse:
			execOpt = cmdShortOption(opt.String())
		case ListOptionSortByTime:
			execOpt = cmdShortOption(opt.String())
		default:
			execOpt = cmdLongOption(opt.String())
		}
		execOpts = append(execOpts, execOpt)
	}
	return execOpts
}
