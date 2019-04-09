package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(clnCmd)
	rootCmd.AddCommand(versionCmd)
}

var rootCmd = &cobra.Command{
	Use:   "migrate-mongo",
	Short: "Migrate mongodb data",
	Long: `A Tool for data migrations between mongodb databases.
Migrations are by streaming way.
Complete documentation is available at https://github.com/Dongss/migrate-mongo`,
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
