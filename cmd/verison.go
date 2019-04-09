package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of migrate-mongo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate-mongo v0.1.0")
	},
}
