package response

type GetProject struct {
	Data struct {
		Attributes struct {
			Name string ``
		} ``
		ID    string ``
		Links struct {
			Self string ``
		} ``
		Relationships struct {
			Streams struct {
				Links struct {
					Related string ``
				} ``
			} ``
		} ``
		Type string ``
	} ``
}
