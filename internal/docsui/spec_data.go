package docsui

// DefaultSpec returns a default DocSpec with example data for the Jolo API documentation.
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
					"Jolo is a logistics and commerce platform that helps businesses/consumers manage deliveries, orders, and operations.",
					"This API powers admin operations, user authentication, and commerce flows like carts and checkout.",
					"Use the sidebar to explore endpoints. Protected endpoints require a Bearer token after login.",
				},
			},
			Onboarding: &UsageSpec{
				Title: "Onboarding & Authentication",
				Notes: []string{
					"To get started, you typically need to authenticate to get a Bearer token. For admin users, use the LOGIN endpoint with your email and password to get a token.",
					"On success, store the token and use it in Authorization: Bearer <token> for protected endpoints. Some endpoints may require additional setup, like 2FA. Follow the steps in the endpoint details.",
                    "Although the documentation is configured to automatically capture any received JWT token and store it in memory for subsequent authenticated requests. You can also manually paste a token in the input field at the top of the page to authenticate.", 
					"Once authenticated, you can explore other endpoints to manage deliveries, orders, and more. you can also use the “Try it out” feature on endpoints to make test API calls directly from the documentation UI.",
                    "The “Try It Out” feature allows you to experiment with the API and see real responses, which is very helpful for understanding how the API works and testing your integration. It’s also smart enough to recognize endpoints that require file attachments—if an endpoint supports file uploads, you’ll see the attachment icon in the upper-left corner; otherwise, it won’t be displayed. Always check the “Content-Type” in the example request. However, keep in mind that any actions performed using this feature will affect your actual data, so use it wisely.",					
                    "Refer to the endpoint details for specific request formats, parameters, and example responses. the documentation is designed to be interactive and user-friendly, so feel free to experiment with the endpoints and see how they work!",
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