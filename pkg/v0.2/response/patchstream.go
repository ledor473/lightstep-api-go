package response

type PatchStream struct {
	Data struct {
		Attributes struct {
			Created_by              string                 `json:"created-by"`
			Created_time            string                 `json:"created-time"`
			CustomData              map[string]interface{} `json:"custom-data,omitempty"`
			Data_last_received_time string                 `json:"data-last-received-time"`
			Name                    string                 `json:"name"`
			Query                   string                 `json:"query"`
		} `json:"attributes"`
		ID    string `json:"id"`
		Links struct {
			Self   string `json:"self"`
			UI     string `json:"ui"`
			UI_p90 string `json:"ui-p90"`
			UI_p95 string `json:"ui-p95"`
			UI_p99 string `json:"ui-p99"`
		} `json:"links"`
		Relationships struct {
			Conditions struct {
				Data []struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"conditions"`
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
