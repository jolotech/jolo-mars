package docsui

// AdminGroup returns a Group with admin-related endpoints for the documentation spec.
func AdminGroup() Group {
	return Group{
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
						Path:    "/admin/auth/login",
						Summary: "Admin Login",
						Auth:    "none",
						Usage: &UsageSpec{
							Title: "Usage",
							Notes: []string{
								"Use this endpoint to authenticate an admin and get a JWT temporaty token based on the admin's account.",
								"e.g. if 2FA is not enabled, you can login with email and password but requires to set up your 2fa.", 
								"If 2FA is enabled, you need to first login with email and password to get a temporary token, then call the 2FA verify endpoint with the code and temporary token to get the access token.",
								"Send email + password in JSON.",
								"On success, store the token and use it in Authorization: Bearer <token> for protected endpoints.",
							},
						},
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
						Path:    "/admin/2fa/setup",
						Summary: "Setup 2FA",
						Auth:    "bearer",
						Responses: []ResponseSpec{
							{Status: 200, Description: "Returns otpauth url", Example: map[string]any{"otpauth_url": "otpauth://totp/..."}},
						},
					},
					{
						ID:      "admin-verify-2fa",
						Method:  "POST",
						Path:    "/admin/2fa/verify",
						Summary: "Verify 2FA code",
						Auth:    "bearer",
						Request: &RequestSpec{
							ContentType: "application/json",
							Example:     map[string]any{"code": "123456"},
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
					{ID: "admin-list", Method: "GET", Path: "/admins", Summary: "List admins", Auth: "bearer"},
					{ID: "admin-delete-all", Method: "DELETE", Path: "/admins", Summary: "Delete all admins", Auth: "bearer"},
				},
			},
		},
	}
}
