package cmd

import (
	"github.com/ErfanMomeniii/data-backupper/internal/config"
	"github.com/ErfanMomeniii/data-backupper/internal/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:                "data back upper",
	Short:              "data back upper",
	PersistentPreRun:   preRun,
	PersistentPostRunE: postRun,
}

func preRun(_ *cobra.Command, _ []string) {
	config.Init()
	err := log.Level.UnmarshalText([]byte(config.C.Logger.Level))
	if err != nil {
		log.Logger.With(zap.Error(err)).Fatal("error in setting log level from config")
	}
}

func postRun(_ *cobra.Command, _ []string) error {
	return log.CloseLogger()
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
