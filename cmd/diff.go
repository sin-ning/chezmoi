package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/twpayne/chezmoi/lib/chezmoi"
	vfs "github.com/twpayne/go-vfs"
)

var diffCmd = &cobra.Command{
	Use:   "diff [targets...]",
	Short: "Write the diff between the target state and the destination state to stdout",
	RunE:  makeRunE(config.runDiffCmd),
}

func init() {
	rootCmd.AddCommand(diffCmd)
}

func (c *Config) runDiffCmd(fs vfs.FS, args []string) error {
	fs = vfs.NewReadOnlyFS(fs)
	mutator := chezmoi.NewLoggingMutator(os.Stdout, chezmoi.NewFSMutator(fs, c.DestDir))
	return c.applyArgs(fs, args, mutator)
}
