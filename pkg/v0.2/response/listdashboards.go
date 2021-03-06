package response

type ListDashboards struct {
	Data []struct {
		Attributes struct {
			Name    string `json:"name"`
			Streams []struct {
				Attributes struct {
					Created_by              string `json:"created-by"`
					Created_time            string `json:"created-time"`
					Data_last_received_time string `json:"data-last-received-time"`
					Name                    string `json:"name"`
					Query                   string `json:"query"`
				} `json:"attributes"`
				ID    string `json:"id"`
				Links struct {
					Self string `json:"self"`
				} `json:"links"`
				Relationships struct {
					Conditions struct {
						Data  interface{} `json:"data"`
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
			} `json:"streams"`
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
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}
