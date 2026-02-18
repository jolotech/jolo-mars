package docsui

func DefaultSpec() DocSpec {
	return DocSpec{
		ProductName: "Jolo API",
		CompanyName: "Jolo",
		Description: "Jolo powers logistics, commerce, and admin operations through secure APIs.",
		BaseURL:     "https://staging.jolojolo.com",
		Version:     "1.0.0",
		QuickStart: QuickStart{
			Title: "Quick Start",
			Steps: []string{
				"Pick an endpoint from the left sidebar.",
				"If it needs auth, paste your Bearer token at the top of the page.",
				"Click an endpoint to view details and optional “Try it out”.",
			},
			Examples: []CodeBlock{
				{
					Title: "Example: Admin Login",
					Lang:  "bash",
					Code: `curl -X POST "$BASE_URL/admin/login" \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@jolo.com","password":"password"}'`,
				},
				{
					Title: "Example: Call protected endpoint",
					Lang:  "bash",
					Code: `curl "$BASE_URL/admin/2fa/setup" \
  -H "Authorization: Bearer YOUR_TOKEN"`,
				},
			},
			Overview: &OverviewSpec{
				Title: "About Jolo & This API",
				Body: []string{
					"Jolo is a logistics and commerce platform that helps businesses manage deliveries, orders, and operations.",
					"This API powers admin operations, user authentication, and commerce flows like carts and checkout.",
					"Use the sidebar to explore endpoints. Protected endpoints require a Bearer token after login.",
				},
			},
		},
		Groups: []Group{
			AdminGroup(),
			UserGroup(),
			LogisticGatewayGroup(),
		},
	}
}



// File: &FileSpec{
//    FieldName: "file",
// 	  Accept: []string{"image/*","video/*","application/pdf"},
//    Multiple: false,
//},