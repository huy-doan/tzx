package command

import (
	"github.com/spf13/cobra"
	application "github.com/test-tzs/nomraeite/batch/application/gmo-aozora/refresh_gmo_aozora_token"
)

var refreshTokenCmd = &cobra.Command{
	Use:   "refresh_gmo_aozora_token",
	Short: "Refresh GMO Aozora Net Bank token",
	Long:  "Refresh GMO Aozora Net Bank token to ensure continuous access to GMO Aozora Net Bank API",
	Run: func(cmd *cobra.Command, args []string) {
		application.Execute()
	},
}

func InitRefreshTokenCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(refreshTokenCmd)
}
