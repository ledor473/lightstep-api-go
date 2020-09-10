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
	"github.com/ledor473/lightstep-api-go/internal/temp/client/dashboards"
	"github.com/ledor473/lightstep-api-go/internal/temp/client/projects"
	"github.com/ledor473/lightstep-api-go/internal/temp/client/services"
	"github.com/ledor473/lightstep-api-go/internal/temp/client/streams"
	"github.com/ledor473/lightstep-api-go/internal/temp/client/workflow_links"
	"github.com/ledor473/lightstep-api-go/internal/temp/models"
	"github.com/spf13/cobra"
)

var (
	jsonResponseCmd = &cobra.Command{
		Use:  "json-response",
		RunE: jsonResponse,
	}

	apikey             string
	organization       string
	project            string
	tempStreamId       string
	tempConditionId    string
	tempDashboardId    string
	tempWorkflowLinkId string
	tempStreamQuery    = "service IN (\"lightstep-api-go\")"
	tempDashboardName  = "Dashboard"
	streamType         = "stream"
	conditionType      = "stream"
	ctx                = context.TODO()
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

	all := []func(*client.LightstepAPI){
		saveProjectsJSON, saveStreamsJSON, saveConditionsJSON, saveServicesJSON,
		saveDashboardsJSON, saveDestinationsJSON, saveWorkflowLinksJSON,
	}
	for _, f := range all {
		f(lsAPI)
	}
	defer cleanupStreams(lsAPI)
	defer cleanupDashboards(lsAPI)
	defer cleanupConditions(lsAPI)

	return nil
}

func saveProjectsJSON(lsAPI *client.LightstepAPI) {
	err := saveJSON("ListProjects", lsAPI, func(c *client.LightstepAPI) (response, error) {
		return c.Projects.ListProjects(ctx, projects.NewListProjectsParams().WithOrganization(organization))
	})
	check(err)
	err = saveJSON("GetProject", lsAPI, func(c *client.LightstepAPI) (response, error) {
		return c.Projects.GetProject(ctx,
			projects.NewGetProjectParams().WithOrganization(organization).WithProject(project))
	})
	check(err)
}

func saveStreamsJSON(lsAPI *client.LightstepAPI) {
	err := saveJSON("ListStreams", lsAPI, func(c *client.LightstepAPI) (response, error) {
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
		check(err)
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
}

func cleanupStreams(lsAPI *client.LightstepAPI) {
	_, err := lsAPI.Streams.DeleteStream(ctx,
		streams.NewDeleteStreamParams().WithOrganization(organization).WithProject(project).WithStreamID(tempStreamId))
	check(err)
}

func saveConditionsJSON(lsAPI *client.LightstepAPI) {
	err := saveJSON("ListConditions", lsAPI, func(c *client.LightstepAPI) (response, error) {
		return c.Conditions.ListConditions(ctx,
			conditions.NewListConditionsParams().WithOrganization(organization).WithProject(project))
	})
	check(err)
	err = saveJSON("PostCondition", lsAPI, func(c *client.LightstepAPI) (response, error) {
		resp, err := c.Conditions.PostCondition(ctx,
			conditions.NewPostConditionParams().WithOrganization(organization).WithProject(project).WithData(
				&models.ConditionRequestBody{
					Data: &models.ConditionRequest{
						Type: &streamType,
						Attributes: &models.ConditionRequestAttributes{
							EvaluationWindowMs: 60000,
							Expression:         "err \u003e 0.1",
							Name:               "Error",
						},
						Relationships: &models.ConditionRequestRelationships{
							Stream: &models.RelatedResourceObject{ID: &tempStreamId, Type: &streamType},
						},
					},
				},
			))
		check(err)
		tempConditionId = resp.Payload.(m)["data"].(m)["id"].(string)
		return resp, err
	})
	check(err)
	err = saveJSON("GetCondition", lsAPI, func(c *client.LightstepAPI) (response, error) {
		return c.Conditions.GetCondition(ctx,
			conditions.NewGetConditionParams().WithOrganization(organization).WithProject(project).WithConditionID(tempConditionId))
	})
	check(err)
	err = saveJSON("PatchCondition", lsAPI, func(c *client.LightstepAPI) (response, error) {
		return c.Conditions.PatchCondition(ctx,
			conditions.NewPatchConditionParams().WithOrganization(organization).WithProject(project).WithConditionID(tempConditionId).WithData(
				&models.ConditionRequestBody{
					Data: &models.ConditionRequest{
						ID:   tempConditionId,
						Type: &streamType,
						Attributes: &models.ConditionRequestAttributes{
							EvaluationWindowMs: 70000,
						},
					},
				},
			))
	})
	check(err)
}

func cleanupConditions(lsAPI *client.LightstepAPI) {
	_, err := lsAPI.Conditions.DeleteCondition(ctx,
		conditions.NewDeleteConditionParams().WithOrganization(organization).WithProject(project).WithConditionID(tempConditionId))
	check(err)
}

func saveServicesJSON(lsAPI *client.LightstepAPI) {
	err := saveJSON("ListServices", lsAPI, func(c *client.LightstepAPI) (response, error) {
		return c.Services.ListServices(ctx, services.NewListServicesParams().WithOrganization(organization).WithProject(project))
	})
	check(err)
}

func saveDashboardsJSON(lsAPI *client.LightstepAPI) {
	err := saveJSON("ListDashboards", lsAPI, func(c *client.LightstepAPI) (response, error) {
		return c.Dashboards.ListDashboards(ctx,
			dashboards.NewListDashboardsParams().WithOrganization(organization).WithProject(project))
	})
	check(err)
	err = saveJSON("CreateDashboard", lsAPI, func(c *client.LightstepAPI) (response, error) {
		resp, err := c.Dashboards.CreateDashboard(ctx,
			dashboards.NewCreateDashboardParams().WithOrganization(organization).WithProject(project).WithData(
				&models.DashboardRequestBody{
					Data: &models.DashboardRequest{
						Type: &streamType,
						Attributes: &models.DashboardAttributes{
							Name:    &tempDashboardName,
							Streams: []*models.StreamResponse{{ID: tempStreamId}},
						},
					},
				},
			))
		check(err)
		tempDashboardId = resp.Payload.(m)["data"].(m)["id"].(string)
		return resp, err
	})
	check(err)
	err = saveJSON("GetDashboard", lsAPI, func(c *client.LightstepAPI) (response, error) {
		return c.Dashboards.GetDashboard(ctx,
			dashboards.NewGetDashboardParams().WithOrganization(organization).WithProject(project).WithDashboardID(tempDashboardId))
	})
	check(err)
	err = saveJSON("PatchDashboard", lsAPI, func(c *client.LightstepAPI) (response, error) {
		return c.Dashboards.PatchDashboard(ctx,
			dashboards.NewPatchDashboardParams().WithOrganization(organization).WithProject(project).WithDashboardID(tempDashboardId).WithData(
				&models.DashboardRequestBody{
					Data: &models.DashboardRequest{
						ID:   tempDashboardId,
						Type: &streamType,
						Attributes: &models.DashboardAttributes{
							Name: &tempDashboardName,
						},
					},
				},
			))
	})
	check(err)
}

func cleanupDashboards(lsAPI *client.LightstepAPI) {
	_, err := lsAPI.Dashboards.DeleteDashboard(ctx,
		dashboards.NewDeleteDashboardParams().WithOrganization(organization).WithProject(project).WithDashboardID(tempDashboardId))
	check(err)
}

func saveDestinationsJSON(lsAPI *client.LightstepAPI) {
	// TODO: Not available in the original Swagger JSON
}

func saveWorkflowLinksJSON(lsAPI *client.LightstepAPI) {
	err := saveJSON("ListWorkflowLinks", lsAPI, func(c *client.LightstepAPI) (response, error) {
		return c.WorkflowLinks.ListWorkflowLinks(ctx,
			workflow_links.NewListWorkflowLinksParams().WithOrganization(organization).WithProject(project))
	})
	check(err)
	rules := make(models.Rules)
	rules["$service"] = []string{"some-service"}
	err = saveJSON("CreateWorkflowLink", lsAPI, func(c *client.LightstepAPI) (response, error) {
		resp, err := c.WorkflowLinks.CreateWorkflowLink(ctx,
			workflow_links.NewCreateWorkflowLinkParams().WithOrganization(organization).WithProject(project).WithData(
				&models.ExternalLinkRequestBody{
					Data: &models.ExternalLinkRequest{
						Type: &streamType,
						Attributes: &models.ExternalLinkAttributes{
							Name:  "WorkflowLink",
							Rules: rules,
							URL:   "https://google.com",
						},
					},
				},
			))
		check(err)
		tempWorkflowLinkId = resp.Payload.(m)["data"].(m)["id"].(string)
		return resp, err
	})
	check(err)
	err = saveJSON("GetWorkflowLink", lsAPI, func(c *client.LightstepAPI) (response, error) {
		return c.WorkflowLinks.GetWorkflowLink(ctx,
			workflow_links.NewGetWorkflowLinkParams().WithOrganization(organization).WithProject(project).WithLinkID(tempWorkflowLinkId))
	})
	check(err)
	err = saveJSON("PatchWorkflowLink", lsAPI, func(c *client.LightstepAPI) (response, error) {
		return c.WorkflowLinks.PatchWorkflowLink(ctx,
			workflow_links.NewPatchWorkflowLinkParams().WithOrganization(organization).WithProject(project).WithLinkID(tempWorkflowLinkId).WithData(
				&models.ExternalLinkRequestBody{
					Data: &models.ExternalLinkRequest{
						Type: &streamType,
						Attributes: &models.ExternalLinkAttributes{
							Name:  "WorkflowLink",
							Rules: rules,
							URL:   "https://google.com",
						},
					},
				},
			))
	})
	check(err)
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
