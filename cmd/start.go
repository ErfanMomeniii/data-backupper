package cmd

import (
	"github.com/ErfanMomeniii/data-backupper/internal/app"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func start() {
	telemetryClearFunc := app.WithTelemetry()

	defer telemetryClearFunc()

	app.WithGracefulShutdown()

	app.Wait()
}
