package cmd

import (
	"fmt"

	"github.com/Dongss/migrate-mongo/mdb"
	"github.com/spf13/cobra"
)

// flags variables
var (
	fSrc     string
	fDst     string
	fIndexes bool
	fShow    bool
)

func init() {
	clnCmd.PersistentFlags().StringVarP(&fSrc, "src", "s", "mongodb://user:pwd@127.0.0.1/database1", "Source mongodb uri (required)")
	clnCmd.PersistentFlags().StringVarP(&fDst, "dst", "d", "mongodb://user:pwd@127.0.0.1/database2", "Destination mongodb uri (required)")
	clnCmd.PersistentFlags().BoolVar(&fIndexes, "index", false, "Include indexes")
	clnCmd.PersistentFlags().BoolVar(&fShow, "show-only", false, "Only show details of source db collection, no migration operation")
	clnCmd.MarkPersistentFlagRequired("src")
	clnCmd.MarkPersistentFlagRequired("dst")
}

var clnCmd = &cobra.Command{
	Use:   "cln <collections> [flags]",
	Short: "Migrate specified collections",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		m := mdb.NewMDB(fSrc, fDst)
		m.Connect()
		defer m.Disconnect()

		co := mdb.ClnOpt{
			IfAll:    false,
			IfIndex:  false,
			ClnNames: args}
		m.Overview(co)
		if fShow == true {
			return
		}
		m.Migrate(co)
		// processing bar
		// fmt.Print("=======\r")
		// time.Sleep(time.Second * 3)
		// fmt.Print("============================\r")
		// time.Sleep(time.Second * 3)
	},
}
