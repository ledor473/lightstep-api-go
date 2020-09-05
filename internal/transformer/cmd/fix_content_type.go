package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(fixContentTypeCmd)
}

var (
	fixContentTypeCmd = &cobra.Command{
		Use:  "fix-content-type",
		RunE: fixContentType,
	}
)

func fixContentType(cmd *cobra.Command, args []string) error {
	doc, err := loads.Spec(input)
	if err != nil {
		return err
	}
	specDoc := doc.Spec()

	if isIncorrectContentType(specDoc.Produces) {
		specDoc.Produces = append(specDoc.Produces, "application/vnd.api+json")
	}

	for path, item := range specDoc.Paths.Paths {
		fmt.Println(path)
		operations := []*spec.Operation{item.Get, item.Put, item.Head, item.Post, item.Delete, item.Options, item.Patch}
		for _, op := range operations {
			if op != nil && isIncorrectContentType(op.Produces) {
				op.Produces = append(op.Produces, "application/vnd.api+json")
			}
		}
	}

	var b []byte
	b, err = json.MarshalIndent(specDoc, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(output, b, 0644)
}

func isIncorrectContentType(produces []string) bool {
	return len(produces) == 1 && produces[0] == "application/json"
}
