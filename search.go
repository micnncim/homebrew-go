package homebrew

import (
	"context"
	"os/exec"
)

type SearchOption string

const (
	SearchOptionCasks    SearchOption = "--casks"
	SearchOptionDesc     SearchOption = "--desc"
	SearchOptionMacports SearchOption = "--macports"
	SearchOptionFink     SearchOption = "--fink"
	SearchOptionOpensuse SearchOption = "--opensuse"
	SearchOptionFedora   SearchOption = "--fedora"
	SearchOptionDebian   SearchOption = "--debian"
	SearchOptionUbuntu   SearchOption = "--ubuntu"
	SearchOptionVerbose  SearchOption = "--verbose"
	SearchOptionDebug    SearchOption = "--debug"
	SearchOptionHelp     SearchOption = "--help"
)

func (o SearchOption) String() string {
	return string(o)
}

func (h *homebrew) Search(ctx context.Context, text string, opts ...SearchOption) error {
	args := make([]string, 0, len(opts)+2)
	args = append(args, searchCmd)
	args = append(args, searchOptions(opts)...)
	args = append(args, text)

	cmd := exec.CommandContext(ctx, rootCmd, args...)
	cmd.Stdout = h.stdout
	cmd.Stderr = h.stderr

	return cmd.Run()
}

func searchOptions(opts []SearchOption) []string {
	execOpts := make([]string, 0, len(opts))
	for _, opt := range opts {
		execOpts = append(execOpts, opt.String())
	}
	return execOpts
}
