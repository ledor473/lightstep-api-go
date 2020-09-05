package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	input  string
	output string
)

var rootCmd = &cobra.Command{
	Use: "transformer",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Nothing to do, use a subcommand instead...")
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&input, "input", "i", "", "Location of the input Swagger file")
	rootCmd.MarkPersistentFlagRequired("input")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Location of the output Swagger file")
	rootCmd.MarkPersistentFlagRequired("output")
}

// Execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
