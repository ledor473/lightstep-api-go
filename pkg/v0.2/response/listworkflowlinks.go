package response

type ListWorkflowLinks struct {
	Data []struct {
		Attributes struct {
			CreatedBy         string `json:"created_by"`
			LastClickedMicros int64  `json:"last_clicked_micros"`
			LastEditedMicros  int64  `json:"last_edited_micros"`
			Name              string `json:"name"`
			Rules             struct {
				Service         []string    `json:"$service"`
				Tm_productCode  interface{} `json:"tm.product_code"`
				UpstreamCluster []string    `json:"upstream_cluster"`
			} `json:"rules"`
			URL string `json:"url"`
		} `json:"attributes"`
		ID    string `json:"id"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Relationships struct {
			Project struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"project"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}
