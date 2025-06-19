package command

import (
	"github.com/spf13/cobra"
	application "github.com/test-tzs/nomraeite/batch/application/gmo-aozora/transfer_status"
)

var transferStatus = &cobra.Command{
	Use:   "gmo_aozora_transfer_status",
	Short: "run gmo_aozora_transfer_status batch job",
	Long:  "run gmo_aozora_transfer_status batch job for checking transfer status through Aozora Bank API",
	Run: func(batch *cobra.Command, args []string) {
		application.Execute()
	},
}

func InitAozoraTransferStatusBatch(rootBatch *cobra.Command) {
	rootBatch.AddCommand(transferStatus)
}
