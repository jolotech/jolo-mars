package docsui

func DefaultSpec() DocSpec {
	return DocSpec{
		ProductName: "Jolo API",
		CompanyName: "Jolo",
		Description: "Jolo powers logistics, commerce, and admin operations through secure APIs.",
		BaseURL:     "/api/v1",
		Version:     "1.0.0",
		QuickStart: QuickStart{
			Title: "Quick Start",
			Steps: []string{
				"1) Pick an endpoint from the left sidebar.",
				"2) If it needs auth, paste your Bearer token at the top of the page.",
				"3) Click an endpoint to view details and optional “Try it out”.",
			},
			Examples: []CodeBlock{
				{
					Title: "Example: Admin Login",
					Lang:  "bash",
					Code: `curl -X POST "$BASE_URL/api/v1/admin/login" \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@jolo.com","password":"password"}'`,
				},
				{
					Title: "Example: Call protected endpoint",
					Lang:  "bash",
					Code: `curl "$BASE_URL/api/v1/admin/2fa/setup" \
  -H "Authorization: Bearer YOUR_TOKEN"`,
				},
			},
		},
		Groups: []Group{
			{
				ID:    "admin",
				Title: "Admin",
				Sections: []Section{
					{
						ID:    "admin-auth",
						Title: "Auth",
						Endpoints: []Endpoint{
							{
								ID:      "admin-login",
								Method:  "POST",
								Path:    "/api/v1/admin/login",
								Summary: "Admin Login",
								Auth:    "none",
								Request: &RequestSpec{
									ContentType: "application/json",
									Example: map[string]any{
										"email":    "dev@jolo.com",
										"password": "password",
									},
								},
								Responses: []ResponseSpec{
									{Status: 200, Description: "Success", Example: map[string]any{"token": "jwt_here"}},
									{Status: 401, Description: "Invalid credentials", Example: map[string]any{"message": "Unauthorized"}},
								},
							},
							{
								ID:      "admin-setup-2fa",
								Method:  "POST",
								Path:    "/api/v1/admin/2fa/setup",
								Summary: "Setup 2FA",
								Auth:    "bearer",
								Responses: []ResponseSpec{
									{Status: 200, Description: "Returns otpauth url", Example: map[string]any{"otpauth_url": "otpauth://totp/..."}},
								},
							},
							{
								ID:      "admin-verify-2fa",
								Method:  "POST",
								Path:    "/api/v1/admin/2fa/verify",
								Summary: "Verify 2FA code",
								Auth:    "bearer",
								Request: &RequestSpec{
									ContentType: "application/json",
									Example: map[string]any{"code": "123456"},
								},
								Responses: []ResponseSpec{
									{Status: 200, Description: "2FA enabled", Example: map[string]any{"enabled": true}},
								},
							},
						},
					},
					{
						ID:    "admin-management",
						Title: "Management",
						Endpoints: []Endpoint{
							{ID: "admin-list", Method: "GET", Path: "/api/v1/admins", Summary: "List admins", Auth: "bearer"},
							{ID: "admin-delete-all", Method: "DELETE", Path: "/api/v1/admins", Summary: "Delete all admins", Auth: "bearer"},
						},
					},
				},
			},
			{
				ID:    "users",
				Title: "Users",
				Sections: []Section{
					{
						ID:    "users-auth",
						Title: "Auth",
						Endpoints: []Endpoint{
							{ID: "user-signup", Method: "POST", Path: "/api/v1/users/signup", Summary: "Signup", Auth: "none"},
							{ID: "user-login", Method: "POST", Path: "/api/v1/users/login", Summary: "Login", Auth: "none"},
							{ID: "user-forgot-password", Method: "POST", Path: "/api/v1/users/forgot-password", Summary: "Forgot Password", Auth: "none"},
						},
					},
				},
			},
			{
				ID:    "carts",
				Title: "Carts",
				Sections: []Section{
					{
						ID:    "carts-main",
						Title: "Cart",
						Endpoints: []Endpoint{
							{ID: "cart-get", Method: "GET", Path: "/api/v1/cart", Summary: "Get cart", Auth: "bearer"},
							{ID: "cart-add", Method: "POST", Path: "/api/v1/cart/items", Summary: "Add item", Auth: "bearer"},
							{ID: "cart-remove", Method: "DELETE", Path: "/api/v1/cart/items/:itemId", Summary: "Remove item", Auth: "bearer"},
						},
					},
				},
			},
		},
	}
}
