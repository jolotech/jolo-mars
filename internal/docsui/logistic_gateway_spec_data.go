package docsui

func LogisticGatewayGroup() Group {
	return Group{
		ID:    "logistics",
		Title: "Logistics Gateway",
		Sections: []Section{
			{
				ID:    "logistics-health",
				Title: "Health",
				Endpoints: []Endpoint{
					{
						ID:      "logistics-health",
						Method:  "GET",
						Path:    "/health",
						Summary: "Service health check",
						Auth:    "none",
					},
				},
			},
		},
	}
}
