package response

type ListConditions struct {
	Data []struct {
		Attributes struct {
			Eval_window_ms int64  ``
			Expression     string ``
			Name           string ``
		} ``
		ID    string ``
		Links struct {
			Self string ``
		} ``
		Relationships struct {
			Project struct {
				Links struct {
					Related string ``
				} ``
			} ``
			Stream struct {
				Links struct {
					Related string ``
				} ``
			} ``
		} ``
		Type string ``
	} ``
	Links struct {
		Self string ``
	} ``
}
