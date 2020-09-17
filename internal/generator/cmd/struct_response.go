package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ChimeraCoder/gojson"
	"github.com/spf13/cobra"
)

var (
	structResponseCmd = &cobra.Command{
		Use:  "struct-response",
		RunE: structResponse,
	}
)

func init() {
	rootCmd.AddCommand(structResponseCmd)
	structResponseCmd.Flags().StringVarP(&input, "input", "i", "", "Input directory")
	structResponseCmd.MarkFlagRequired("input")
}

func structResponse(cmd *cobra.Command, args []string) error {
	err := filepath.Walk(input, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		structName := strings.TrimSuffix(filepath.Base(file.Name()), filepath.Ext(file.Name()))
		fileName := strings.ToLower(structName)
		generated, err := gojson.Generate(file, gojson.ParseJson, structName, "response", []string{"json"}, false, true)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(filepath.Join(output, fmt.Sprintf("%s.go", fileName)), generated, 0644)
	})
	check(err)
	return nil
}
