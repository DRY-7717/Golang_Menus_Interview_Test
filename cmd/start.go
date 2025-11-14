package cmd

import (
	"golang_menu_interview/internal/app"

	"github.com/spf13/cobra"
)

// ini start cmd adalah subcommand dari rootCmd
// karna ini subcommand dari rootCmd nanti saat dijalankan nya tinggal core-api start
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start",
	Long:  "start",
	Run: func(cmd *cobra.Command, args []string) {
		// Call function root api
		app.RunServer()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
