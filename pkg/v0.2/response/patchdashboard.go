package response

type PatchDashboard struct {
	Data struct {
		Attributes struct {
			Name    string      `json:"name"`
			Streams interface{} `json:"streams"`
		} `json:"attributes"`
		ID    string `json:"id"`
		Links struct {
			Self string `json:"self"`
			UI   string `json:"ui"`
		} `json:"links"`
		Relationships struct {
			Project struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"project"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}
