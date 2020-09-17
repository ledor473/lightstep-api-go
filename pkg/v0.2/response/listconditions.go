package response

type ListConditions struct {
	Data []struct {
		Attributes struct {
			Eval_window_ms int64  `json:"eval-window-ms"`
			Expression     string `json:"expression"`
			Name           string `json:"name"`
		} `json:"attributes"`
		ID    string `json:"id"`
		Links struct {
			Self string `json:"self"`
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
			Stream struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"stream"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}
