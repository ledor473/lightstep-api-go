package response

type GetStream struct {
	Data struct {
		Attributes struct {
			Created_by              string ``
			Created_time            string ``
			Data_last_received_time string ``
			Name                    string ``
			Query                   string ``
		} ``
		ID    string ``
		Links struct {
			Self   string ``
			UI     string ``
			UI_p90 string ``
			UI_p95 string ``
			UI_p99 string ``
		} ``
		Relationships struct {
			Conditions struct {
				Links struct {
					Related string ``
				} ``
			} ``
			Project struct {
				Links struct {
					Related string ``
				} ``
			} ``
		} ``
		Type string ``
	} ``
}
