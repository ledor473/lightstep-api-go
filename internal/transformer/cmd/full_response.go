package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-openapi/jsonreference"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(fullResponseCmd)
	fullResponseCmd.Flags().StringVarP(&responsePackage, "package", "p", "", "Location to the response package")
	fullResponseCmd.MarkFlagRequired("package")
}

var (
	fullResponseCmd = &cobra.Command{
		Use:  "full-response",
		RunE: fullResponse,
	}
	responsePackage string
	structs         []string
)

func fullResponse(cmd *cobra.Command, args []string) error {
	doc, err := loads.Spec(input)
	if err != nil {
		return err
	}
	specDoc := doc.Spec()

	for path, item := range specDoc.Paths.Paths {
		fmt.Println(path)
		operations := []*spec.Operation{item.Get, item.Put, item.Head, item.Post, item.Delete, item.Options, item.Patch}
		for _, op := range operations {
			if op != nil && responseStructAvailable(op.ID) {
				for code, resp := range op.Responses.StatusCodeResponses {
					if code/200 == 1 {
						t := strings.Title(op.ID)
						specDoc.Definitions[t] = spec.Schema{
							VendorExtensible: spec.VendorExtensible{Extensions: createTypeExtension(t)},
						}

						resp.Schema = &spec.Schema{
							SchemaProps: spec.SchemaProps{
								Ref: spec.Ref{
									jsonreference.MustCreateRef(fmt.Sprintf("#/definitions/%s", t)),
								},
							},
						}

						op.Responses.StatusCodeResponses[code] = resp
					}
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

func createTypeExtension(t string) map[string]interface{} {
	var ext = make(map[string]interface{})
	ext["x-go-type"] = make(map[string]interface{})
	ext["x-go-type"].(map[string]interface{})["type"] = t
	var imp = make(map[string]string)
	imp["package"] = "github.com/ledor473/lightstep-api-go/pkg/v0.2/response"
	ext["x-go-type"].(map[string]interface{})["import"] = imp

	return ext
}

func responseStructAvailable(id string) bool {
	if len(structs) == 0 {
		findAvailableStructs()
	}
	for _, s := range structs {
		if strings.EqualFold(s, id) {
			return true
		}
	}
	return false
}

func findAvailableStructs() {
	filepath.Walk(responsePackage, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			structs = append(structs, strings.TrimSuffix(info.Name(), filepath.Ext(info.Name())))
		}
		return nil
	})
}
