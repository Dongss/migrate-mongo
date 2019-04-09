package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// flags variables
var (
	fSrc     string
	fDst     string
	fIndexes bool
)

func init() {
	clnCmd.PersistentFlags().StringVarP(&fSrc, "src", "s", "mongodb://user:pwd@127.0.0.1/database1", "Source mongodb uri (required)")
	clnCmd.PersistentFlags().StringVarP(&fDst, "dst", "d", "mongodb://user:pwd@127.0.0.1/database2", "Destination mongodb uri (required)")
	clnCmd.PersistentFlags().BoolVar(&fIndexes, "index", false, "Include indexes")
	clnCmd.MarkPersistentFlagRequired("src")
	clnCmd.MarkPersistentFlagRequired("dst")
}

var clnCmd = &cobra.Command{
	Use:   "cln <collections> [flags]",
	Short: "Migrate specified collections",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("not finised yet")
	},
}
