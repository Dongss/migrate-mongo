package cmd

import (
	"errors"
	"fmt"

	"github.com/Dongss/migrate-mongo/mdb"
	"github.com/spf13/cobra"
)

// flags variables
var (
	fSrc      string
	fDst      string
	fInterval int64
	fIndexes  bool
	fShow     bool
	fAll      bool
	fBatch    int32
)

func init() {
	clnCmd.PersistentFlags().StringVarP(&fSrc, "src", "s", "mongodb://user:pwd@127.0.0.1/database1", "Source mongodb uri (required)")
	clnCmd.PersistentFlags().StringVarP(&fDst, "dst", "d", "mongodb://user:pwd@127.0.0.1/database2", "Destination mongodb uri (required)")
	clnCmd.PersistentFlags().BoolVar(&fIndexes, "index", false, "Include indexes, create indexes before inserting data")
	clnCmd.PersistentFlags().Int64VarP(&fInterval, "interval", "i", 0, "Interval of each single insert, milliseconds")
	clnCmd.PersistentFlags().BoolVar(&fShow, "show-only", false, "Only show details of source db collection, no migration operation")
	clnCmd.PersistentFlags().BoolVar(&fAll, "all", false, "Migrate all collections")
	clnCmd.PersistentFlags().Int32VarP(&fBatch, "batch", "b", 1, "Batch insert, count of each inserting")
	clnCmd.MarkPersistentFlagRequired("src")
	clnCmd.MarkPersistentFlagRequired("dst")
}

var clnCmd = &cobra.Command{
	Use:   "cln <collections> [flags]",
	Short: "Migrate specified collections",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 && fAll != true {
			return errors.New("requires collection names")
		}
		if len(args) > 0 && fAll == true {
			return errors.New("migrate all, should not pass collection names")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		m := mdb.NewMDB(fSrc, fDst)
		m.Connect()
		defer m.Disconnect()

		co := mdb.ClnOpt{
			IfAll:    fAll,
			ClnNames: args}
		m.Overview(co)
		if fShow == true {
			return
		}
		m.Migrate(co, mdb.MigOpt{
			Interval: fInterval,
			IfIndex:  false,
			FBatch:   fBatch,
		})
	},
}
