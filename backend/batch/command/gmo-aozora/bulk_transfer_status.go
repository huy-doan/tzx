package command

import (
	"github.com/spf13/cobra"
	application "github.com/test-tzs/nomraeite/batch/application/gmo-aozora/bulk_transfer_status"
)

var bulkTransferStatus = &cobra.Command{
	Use:   "gmo_aozora_bulk_transfer_status",
	Short: "run gmo_aozora_bulk_transfer_status batch job",
	Long:  "run gmo_aozora_bulk_transfer_status batch job for checking transfer status through Aozora Bank API",
	Run: func(batch *cobra.Command, args []string) {
		application.Execute()
	},
}

func InitAozoraBulkTransferStatusBatch(rootBatch *cobra.Command) {
	rootBatch.AddCommand(bulkTransferStatus)
}
