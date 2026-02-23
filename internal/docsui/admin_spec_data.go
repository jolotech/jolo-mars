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
								"email": "isihaqabdullahi01+1@gmail.com",
								"password": "med#*Jm@sPa@sRxE",
							},
						},
						Responses: []ResponseSpec{
							{Status: 200, Description: "2fa Success", Example: map[string]any{
								"status": "success",
								"message": "2FA not setup. Please use setup endpoint",
								"data": map[string]any{
									"requires_2fa": true,
									"requires_2fa_message": "2FA not setup for this account, please setup 2FA to secure your account",
									"password_change_required": true,
									"setup_token": "JWT token e.g -- eyJhbGciOiJIUzI1NiIsIn",
								},
								"code": 200,
							}},
							{Status: 401, Description: "Invalid credentials", Example: map[string]any{"status": "error","message": "invalid credentials", "code": 401,}},
						},
					},
					{
						ID:      "admin-setup-2fa",
						Method:  "GET",
						Path:    "/admin/auth/2fa/setup",
						Summary: "Setup 2FA",
						Auth:    "bearer",
						Usage: &UsageSpec{
							Title: "Usage",
							Notes: []string{
								"Use this endpoint to setup 2FA for the admin account.",
								"Call this endpoint after successful login to get the otpauth url to generate QR code for 2FA setup.",
								"The otpauth url can be used to generate a QR code with reactJS library (e.g. using qrcode.react) or use a web tool (e.g. using https://qr.io/).",
								"Scan the QR code with your 2FA app (e.g. Google Authenticator) to get the 6-digit code for verification.",
								"Requires Authorization header with Bearer setup_token from login response.",
							},
						},
						Responses: []ResponseSpec{
							{
								Status: 200, 
								Description: "Returns otpauth url for generating QR code ",
							    Example: map[string]any{"status": "success", "message": "2fa setup successful", 
								    "data": map[string]any{
									    "otpauth_url": "otpauth://totp/Jolo%20....",
								    },
								    "code": 200,
							    },
							},
							{Status: 401, Description: "Invalide auth token", Example: map[string]any{"status": "error", "message": "Invalid token purpose", "code": 401}},
					    },
					},
					{
						ID:      "admin-verify-2fa",
						Method:  "POST",
						Path:    "/admin/auth/2fa/confirm",
						Summary: "Verify 2FA code",
						Auth:    "bearer",
						Usage: &UsageSpec{
							Title: "Usage",
							Notes: []string{
								"Use this endpoint to verify the 2FA code and enable 2FA for the admin account.",
								"Call this endpoint after getting the setup token from login and setting up 2FA with the otpauth url.",
								"Requires Authorization header with Bearer setup_token or two_fa_token from login response.",
								"On success, 2FA will be enabled for the account and you can use the same code to login next time.",
								"After enabling 2FA, you need to use the 2FA token from login response to verify the code.",
								"If password change is required, you need to change the password first before you can use the access_token for protected endpoints.",
								"On failure, 2FA will not be enabled and you can retry with the correct code.",
							    "If 2FA is already enabled, you can call this endpoint to verify the code and get a new access token without needing to set up 2FA again, just make sure to use the two_fa_token from login response for verification instead of the setup_token.",
								"",
								"NOTE:   The setup_token and two_fa_token from login response can both be used to call this endpoint for verification, but the two_fa_token is meant to last for 15 minutes and can only be used for verification, while the setup_token is meant to be used for setup and verification and may last longer (e.g. 30 minutes) to allow for setup and verification process.",
							},
						},
						Request: &RequestSpec{
							ContentType: "application/json",
							Example:     map[string]any{"code": "123456"},
						},
						Responses: []ResponseSpec{
							{
								Status: 200, 
								Description: "Successfully enabled 2FA response ", 
								Example: map[string]any{
									"status": "success",
								    "message": "password_change_required",
								    "data": map[string]any{
									   "password_change_required": true,
									   "setup_token": "JWT token e.g eyJhbGciOiJIUzI1....",
								    },
								    "code": 200,
								},
							},
							{Status: 400, Description: "Invalid code", Example: map[string]any{"status": "error", "message": "invalid 2FA code", "code": 400}},
						},
					},
				},
			},
			{
				ID:    "admin-dashboard",
				Title: "Dashboard",
				Endpoints: []Endpoint{
					{
						ID: "admin-change-password", 
						Method: "PUT", 
						Path: "/admin/dash/change-password", 
						Summary: "Change password", 
						Auth: "bearer",
						Usage: &UsageSpec{
							Title: "Usage",
							Notes: []string{
								"Use this endpoint to change the password of the admin account.",
								"Requires Authorization header with Bearer token from 2fa verification response.",
								"If password change is required after login, you need to change the password before you can use the access token for protected endpoints.",
								"You can only change admin password if the password_change_required field in the login or 2fa verification response is true, otherwise you will get an access_token from the 2fa verification response and can use it to access protected endpoints without needing to change password.",
								"Send new_password, current_password and confirm_password in JSON.",
								"On success, the password will be updated and you can use the new access token in the response to access protected endpoints.",
								"On failure, the password will not be changed and you can retry with the correct current password and valid new password that meets the password requirements.",
							},
						},
						Request: &RequestSpec{
							ContentType: "application/json",
							Example:     map[string]any{"new_password": "wg+wlHvnow9", "current_password": "4k2$L&jU3bA982vy","confirm_password": "wg+wlHvnow9"},
						},
						Responses: []ResponseSpec{
							{
								Status: 400, 
								Description: "Password Change error", 
								Example: map[string]any{
								"status": "error",
								"message": "new password must include at least one uppercase letter",
								"code": 400,
								},
							},
							{
								Status: 200,
								Description: "Password Change success", 
								Example: map[string]any{
									"status": "success",
									"message": "password updated successfully",
									"data": map[string]any{
										"access_token": "JWT token e.g eyJhbGciOiJIUzI1NiI",
                                        "password_change_required": false,
                                        "admin": map[string]any{
											"public_id": "AMbm48jjMPaTz0Q",
											"name": "Jolo Mars",
											"email": "admin@jolo.com",
											"role": "admin",
											"createdAt": "2026-02-08T06:12:44.069+01:00",
											"updatedAt": "2026-02-08T06:12:44.069+01:00",
										},
									},
									"code": 200,
								},
							},
						},
					},
				},
			},
			{
				ID:    "admin-management",
				Title: "Management - (pending)",
				Endpoints: []Endpoint{
					{ID: "admin-list", Method: "GET", Path: "/admins", Summary: "List admins", Auth: "bearer"},
					{ID: "admin-delete-all", Method: "DELETE", Path: "/admins", Summary: "Delete all admins", Auth: "bearer"},
				},
			},
		},
	}
}
