package command

import (
	application "github.com/makeshop-jp/master-console/batch/application/gmo-aozora/transfer_request"
	"github.com/spf13/cobra"
)

var transferRequest = &cobra.Command{
	Use:   "gmo_aozora_transfer_request",
	Short: "run gmo_aozora_transfer_request batch job",
	Long:  "run gmo_aozora_transfer_request batch job for transferring money through Aozora Bank API",
	Run: func(batch *cobra.Command, args []string) {
		application.Execute()
	},
}

func InitAozoraTransferRequestBatch(rootBatch *cobra.Command) {
	rootBatch.AddCommand(transferRequest)
}
