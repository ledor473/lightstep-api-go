package response

type ListServices struct {
	Data struct {
		Items []struct {
			Attributes struct {
				LastSeen string `json:"last_seen"`
				Name     string `json:"name"`
			} `json:"attributes"`
			Relationships struct {
				Project struct {
					Links struct {
						Related string `json:"related"`
					} `json:"links"`
				} `json:"project"`
			} `json:"relationships"`
		} `json:"items"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Type string `json:"type"`
	} `json:"data"`
}
