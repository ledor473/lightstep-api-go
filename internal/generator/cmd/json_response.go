package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	openapi "github.com/go-openapi/runtime/client"
	"github.com/ledor473/lightstep-api-go/internal/temp/client"
	"github.com/ledor473/lightstep-api-go/internal/temp/client/conditions"
	"github.com/ledor473/lightstep-api-go/internal/temp/client/projects"
	"github.com/ledor473/lightstep-api-go/internal/temp/client/streams"
	"github.com/ledor473/lightstep-api-go/internal/temp/models"
	"github.com/spf13/cobra"
)

var (
	jsonResponseCmd = &cobra.Command{
		Use:  "json-response",
		RunE: jsonResponse,
	}

	apikey          string
	organization    string
	project         string
	tempStreamId    string
	tempStreamQuery = "service IN (\"lightstep-api-go\")"
	streamType      = "stream"
	conditionType   = "stream"
)

func init() {
	rootCmd.AddCommand(jsonResponseCmd)
	jsonResponseCmd.Flags().StringVar(&apikey, "apikey", "", "API Key for LightStep")
	jsonResponseCmd.MarkFlagRequired("apikey")
	jsonResponseCmd.Flags().StringVar(&organization, "org", "", "LightStep Organization")
	jsonResponseCmd.MarkFlagRequired("organization")
	jsonResponseCmd.Flags().StringVarP(&project, "project", "p", "", "LightStep Project")
	jsonResponseCmd.MarkFlagRequired("project")
}

type clientCall = func(*client.LightstepAPI) (response, error)
type response interface {
	GetPayload() interface{}
}
type m = map[string]interface{}

func jsonResponse(cmd *cobra.Command, args []string) error {
	lsAPI := client.New(client.Config{AuthInfo: openapi.APIKeyAuth("Authorization", "header", apikey)})

	ctx := context.TODO()

	// Projects
	err := saveJSON("ListProjects", lsAPI, func(c *client.LightstepAPI) (response, error) {
		return c.Projects.ListProjects(ctx, projects.NewListProjectsParams().WithOrganization(organization))
	})
	check(err)
	err = saveJSON("GetProject", lsAPI, func(c *client.LightstepAPI) (response, error) {
		return c.Projects.GetProject(ctx,
			projects.NewGetProjectParams().WithOrganization(organization).WithProject(project))
	})
	check(err)

	// Streams
	err = saveJSON("ListStreams", lsAPI, func(c *client.LightstepAPI) (response, error) {
		return c.Streams.ListStreams(ctx,
			streams.NewListStreamsParams().WithOrganization(organization).WithProject(project))
	})
	check(err)
	err = saveJSON("PostStream", lsAPI, func(c *client.LightstepAPI) (response, error) {
		resp, err := c.Streams.PostStream(ctx,
			streams.NewPostStreamParams().WithOrganization(organization).WithProject(project).WithData(
				&models.CreateOrUpdateBody{
					Data: &models.CreateOrUpdateRequest{
						Type: &streamType,
						Attributes: &models.StreamRequestAttributes{
							Name:  &tempStreamQuery,
							Query: tempStreamQuery,
						},
					},
				},
			))
		tempStreamId = resp.Payload.(m)["data"].(m)["id"].(string)
		return resp, err
	})
	check(err)
	err = saveJSON("GetStream", lsAPI, func(c *client.LightstepAPI) (response, error) {
		return c.Streams.GetStream(ctx,
			streams.NewGetStreamParams().WithOrganization(organization).WithProject(project).WithStreamID(tempStreamId))
	})
	check(err)
	err = saveJSON("PatchStream", lsAPI, func(c *client.LightstepAPI) (response, error) {
		return c.Streams.PatchStream(ctx,
			streams.NewPatchStreamParams().WithOrganization(organization).WithProject(project).WithStreamID(tempStreamId).WithData(
				&models.CreateOrUpdateBody{
					Data: &models.CreateOrUpdateRequest{
						Type: &streamType,
						Attributes: &models.StreamRequestAttributes{
							Name: &tempStreamQuery,
						},
					},
				},
			))
	})
	check(err)

	defer func() { // Cleanup at the end because we still need it
		_, err = lsAPI.Streams.DeleteStream(ctx,
			streams.NewDeleteStreamParams().WithOrganization(organization).WithProject(project).WithStreamID(tempStreamId))
		check(err)
	}()

	// Conditions
	err = saveJSON("ListConditions", lsAPI, func(c *client.LightstepAPI) (response, error) {
		return c.Conditions.ListConditions(ctx,
			conditions.NewListConditionsParams().WithOrganization(organization).WithProject(project))
	})
	check(err)
	// TODO: Everything below

	// Dashboards

	// Destinations

	// Services

	// WorkflowLinks

	return nil
}

func saveJSON(operation string, client *client.LightstepAPI, f clientCall) error {
	resp, err := f(client)
	if err != nil {
		return err
	}

	var b []byte
	b, err = json.MarshalIndent(resp.GetPayload(), "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath.Join(output, fmt.Sprintf("%s.json", operation)), b, 0644)
}
