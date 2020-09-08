package response

type GetProject struct {
	Data struct {
		Attributes struct {
			Name string `json:"name"`
		} `json:"attributes"`
		ID    string `json:"id"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Relationships struct {
			Streams struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"streams"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}
