package homebrew

import (
	"context"
	"os/exec"
)

type InstallOption string

const (
	InstallOptionDebug              InstallOption = "debug"
	InstallOptionEnv                InstallOption = "env"
	InstallOptionIgnoreDependencies InstallOption = "ignore-dependencies"
	InstallOptionOnlyDependencies   InstallOption = "only-dependencies"
	InstallOptionCC                 InstallOption = "cc"
	InstallOptionBuildFromSource    InstallOption = "build-from-source"
	InstallOptionForceBottle        InstallOption = "force-bottle"
	InstallOptionIncludeTest        InstallOption = "include-test"
	InstallOptionDevel              InstallOption = "devel"
	InstallOptionHEAD               InstallOption = "HEAD"
	InstallOptionFetchHEAD          InstallOption = "fetch-HEAD"
	InstallOptionKeepTmp            InstallOption = "keep-tmp"
	InstallOptionBuildBottle        InstallOption = "build-bottle"
	InstallOptionBottleArch         InstallOption = "bottle-arch"
	InstallOptionForce              InstallOption = "force"
	InstallOptionVerbose            InstallOption = "verbose"
	InstallOptionDisplayTimes       InstallOption = "display-times"
	InstallOptionInteractive        InstallOption = "interactive"
	InstallOptionGit                InstallOption = "git"
	InstallOptionHelp               InstallOption = "help"
)

func (o InstallOption) String() string {
	return string(o)
}

func (h *homebrew) Install(ctx context.Context, formula string, opts ...InstallOption) error {
	args := append([]string{installCmd}, installOptions(opts)...)
	args = append(args, formula)

	cmd := exec.CommandContext(ctx, rootCmd, args...)
	cmd.Stdout = h.stdout
	cmd.Stderr = h.stderr

	return cmd.Run()
}

func installOptions(opts []InstallOption) []string {
	execOpts := make([]string, 0, len(opts))
	for _, opt := range opts {
		execOpts = append(execOpts, cmdLongOption(opt.String()))
	}
	return execOpts
}
