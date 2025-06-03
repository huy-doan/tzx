package batch

import (
	"os"

	"github.com/spf13/cobra"
	command "github.com/test-tzs/nomraeite/batch/command/gmo-aozora"
)

// rootbatch represents the base command when called without any subcommands
var rootBatch = &cobra.Command{
	Use:   "backend-job",
	Short: "backend-job is a job runner",
	Long:  `backend-job is a job runner`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootbatch.
func Execute() {
	err := rootBatch.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	command.InitRefreshTokenCommand(rootBatch)
	command.InitAozoraTransferRequestBatch(rootBatch)
}
