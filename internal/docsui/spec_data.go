// package docsui

// func DefaultSpec() DocSpec {
// 	return DocSpec{
// 		ProductName: "Jolo API",
// 		CompanyName: "Jolo",
// 		Description: "Jolo powers logistics, commerce, and admin operations through secure APIs.",
// 		BaseURL:     "https://staging.jolojolo.com/v1",
// 		Version:     "1.0.0",
// 		QuickStart: QuickStart{
// 			Title: "Quick Start",
// 			Steps: []string{
// 				"Pick an endpoint from the left sidebar.",
// 				"If it needs auth, paste your Bearer token at the top of the page.",
// 				"Click an endpoint to view details and optional “Try it out”.",
// 			},
// 			Examples: []CodeBlock{
// 				{
// 					Title: "Example: Admin Login",
// 					Lang:  "bash",
// 					Code: `curl -X POST "$BASE_URL/admin/login" \
//         -H "Content-Type: application/json" \
//         -d '{"email":"admin@jolo.com","password":"password"}'`,
// 				},
// 				{
// 					Title: "Example: Call protected endpoint",
// 					Lang:  "bash",
// 					Code: `curl "$BASE_URL/v1/admin/2fa/setup" \
//         -H "Authorization: Bearer YOUR_TOKEN"`,
// 				},
// 			},
// 			Overview: &OverviewSpec{
// 				Title: "About Jolo & This API",
// 				Body: []string{
// 					"Jolo is a logistics and commerce platform that helps businesses manage deliveries, orders, and operations.",
//                     "This API powers admin operations, user authentication, and commerce flows like carts and checkout.",
//                     "Use the sidebar to explore endpoints. Protected endpoints require a Bearer token after login.",
// 				},
// 			},
// 		},
// 		Groups: []Group{
// // ======== ADMIN GROUP WITH NESTED SECTIONS (AUTH + MANAGEMENT etc..) ========
// 			{
// 				ID:    "admin",
// 				Title: "Admin",
// 				Sections: []Section{
// 					{
// 						ID:    "admin-auth",
// 						Title: "Auth",
// 						Endpoints: []Endpoint{
// 							{
// 								ID:      "admin-login",
// 								Method:  "POST",
// 								Path:    "/admin/login",
// 								Summary: "Admin Login",
// 								Auth:    "none",
// 								Usage: &UsageSpec{
// 								    Title: "Usage",
//                                     Notes: []string{
// 										"Use this endpoint to authenticate an admin and get an access token.",
//                                         "Send email + password in JSON.",
//                                         "On success, store the token and use it in Authorization: Bearer <token> for protected endpoints.",
//                                     },
// 								},
// 								Request: &RequestSpec{
// 									ContentType: "multipart/form-data",
// 									Example: map[string]any{
// 										"email":    "dev@jolo.com",
// 										"password": "password",
// 									},
// 									File: &FileSpec{
// 										FieldName: "file",
// 										Accept: []string{"image/*","video/*","application/pdf"},
// 										Multiple: false,
// 									},
// 								},
// 								Responses: []ResponseSpec{
// 									{Status: 200, Description: "Success", Example: map[string]any{"token": "jwt_here"}},
// 									{Status: 401, Description: "Invalid credentials", Example: map[string]any{"message": "Unauthorized"}},
// 								},
// 							},
// 							{
// 								ID:      "admin-setup-2fa",
// 								Method:  "POST",
// 								Path:    "/api/v1/admin/2fa/setup",
// 								Summary: "Setup 2FA",
// 								Auth:    "bearer",
// 								Responses: []ResponseSpec{
// 									{Status: 200, Description: "Returns otpauth url", Example: map[string]any{"otpauth_url": "otpauth://totp/..."}},
// 								},
// 							},
// 							{
// 								ID:      "admin-verify-2fa",
// 								Method:  "POST",
// 								Path:    "/api/v1/admin/2fa/verify",
// 								Summary: "Verify 2FA code",
// 								Auth:    "bearer",
// 								Request: &RequestSpec{
// 									ContentType: "application/json",
// 									Example: map[string]any{"code": "123456"},
// 								},
// 								Responses: []ResponseSpec{
// 									{Status: 200, Description: "2FA enabled", Example: map[string]any{"enabled": true}},
// 								},
// 							},
// 						},
// 					},
// 		// ======== MANAGEMENT SECTION NESTED UNDER ADMIN ========
// 					{
// 						ID:    "admin-management",
// 						Title: "Management",
// 						Endpoints: []Endpoint{
// 							{ID: "admin-list", Method: "GET", Path: "/api/v1/admins", Summary: "List admins", Auth: "bearer"},
// 							{ID: "admin-delete-all", Method: "DELETE", Path: "/api/v1/admins", Summary: "Delete all admins", Auth: "bearer"},
// 						},
// 					},
// 				},
// 			},

// 	// ======== USERS GROUP WITH NESTED SECTIONS (AUTH + CART etc..) ========
// 			{
//                 ID: "users",
//                 Title: "Users",
//                 Sections: []Section{
// 					{
// 						ID: "users-auth",
//                         Title: "Auth",
//                         Endpoints: []Endpoint{
// 							{ ID:"user-login", Method:"POST", Path:"/users/auth/login", Summary:"Login", Auth:"none" },
//                            { ID:"user-forgot-password", Method:"POST", Path:"/users/auth/forgot-password", Summary:"Forgot Password", Auth:"none" },
//                         },
//                     },
//           // ======== CART SECTION NESTED UNDER USERS ========
// 	                {
// 						ID: "users-cart",
//                         Title: "Cart",
//                          Endpoints: []Endpoint{
//                             { ID:"cart-get", Method:"GET", Path:"/users/cart", Summary:"Get Cart", Auth:"bearer" },
//                             { ID:"cart-add", Method:"POST", Path:"/users/cart/items", Summary:"Add Item", Auth:"bearer" },
//                         },
//                     },
//                 },
//              },
// 		},
// 	}
// }


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
