package cmd

import (
	"bufio"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var (
	tweakResponseCmd = &cobra.Command{
		Use:  "tweak-response",
		RunE: tweakResponse,
	}
)

func init() {
	rootCmd.AddCommand(tweakResponseCmd)
	tweakResponseCmd.Flags().StringVarP(&input, "input", "i", "", "Input directory")
	tweakResponseCmd.MarkFlagRequired("input")
}

func tweakResponse(cmd *cobra.Command, args []string) error {
	// TODO: Could be fancier by using go/ast + go/parser

	err := filepath.Walk(input, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if strings.Contains(path, "workflowlink") {
			data, err := tweakWorkflowLink(path)
			check(err)
			return ioutil.WriteFile(path, data, 0644)
		}

		if strings.Contains(path, "stream") {
			data, err := tweakStream(path)
			check(err)
			return ioutil.WriteFile(path, data, 0644)
		}

		return nil
	})
	check(err)
	return nil
}

func tweakWorkflowLink(path string) ([]byte, error) {
	file, err := os.Open(path)
	defer file.Close()
	check(err)

	lines := replaceDef(file, "Rules *struct", "} *`json:\"rules\"`", rulesDef)

	return format.Source([]byte(strings.Join(lines, "\n")))
}

func tweakStream(path string) ([]byte, error) {
	file, err := os.Open(path)
	defer file.Close()
	check(err)

	lines := replaceDef(file, "Custom_data *struct", "} *`json:\"custom-data\"`", customDataDef)

	return format.Source([]byte(strings.Join(lines, "\n")))
}

func replaceDef(f *os.File, start, end, replacement string) []string {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string

	var found = false
	for scanner.Scan() {
		l := scanner.Text()
		if !found {
			if m, _ := regexp.Match(start, []byte(l)); m {
				found = true
			} else {
				lines = append(lines, l)
			}
		} else {
			if m, _ := regexp.Match(end, []byte(l)); m {
				found = false
				lines = append(lines, replacement)
			}
		}
	}
	return lines
}

const rulesDef = "Rules map[string][]string `json:\"rules\"`"
const customDataDef = "CustomData map[string]interface{} `json:\"custom-data,omitempty\"`"
