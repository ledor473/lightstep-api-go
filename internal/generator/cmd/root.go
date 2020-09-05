package cmd

import (
	"fmt"
	"os"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
)

var (
	output  string
	rootCmd = &cobra.Command{
		Use: "generator",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Nothing to do, use a subcommand instead...")
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Location of the output files")
	rootCmd.MarkPersistentFlagRequired("output")
}

// Execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func check(err error) {
	if err != nil {
		pp.Print(err)
		os.Exit(1)
	}
}
