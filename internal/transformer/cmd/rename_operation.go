package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(renameOperationCmd)
}

var (
	renameOperationCmd = &cobra.Command{
		Use:  "rename-operation",
		RunE: renameOperation,
	}
)

func renameOperation(cmd *cobra.Command, args []string) error {
	doc, err := loads.Spec(input)
	if err != nil {
		return err
	}
	specDoc := doc.Spec()

	for path, item := range specDoc.Paths.Paths {
		fmt.Println(path)
		operations := []*spec.Operation{item.Get, item.Put, item.Head, item.Post, item.Delete, item.Options, item.Patch}
		for _, op := range operations {
			if op != nil && strings.HasSuffix(op.ID, "ID") {
				op.ID = strings.TrimSuffix(op.ID, "ID")
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
