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
	rootCmd.AddCommand(basicResponseCmd)
}

var (
	basicResponseCmd = &cobra.Command{
		Use:  "basic-response",
		RunE: basicResponse,
	}
)

func basicResponse(cmd *cobra.Command, args []string) error {
	doc, err := loads.Spec(input)
	if err != nil {
		return err
	}
	specDoc := doc.Spec()

	for path, item := range specDoc.Paths.Paths {
		fmt.Println(path)
		operations := []*spec.Operation{item.Get, item.Put, item.Head, item.Post, item.Delete, item.Options, item.Patch}
		for _, op := range operations {
			if op != nil {
				for code, resp := range op.Responses.StatusCodeResponses {
					resp.Schema = &spec.Schema{
						SchemaProps: spec.SchemaProps{
							Type:                 spec.StringOrArray{"object"},
							AdditionalProperties: &spec.SchemaOrBool{Allows: true},
						},
					}

					op.Responses.StatusCodeResponses[code] = resp
				}
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
