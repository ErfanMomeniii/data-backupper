package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:                "data back upper",
	Short:              "data back upper",
	PersistentPreRun:   preRun,
	PersistentPostRunE: postRun,
}

func preRun(_ *cobra.Command, _ []string) {
}

func postRun(_ *cobra.Command, _ []string) error {
	return nil
}

func init() {
}

func Execute() error {
	return rootCmd.Execute()
}
