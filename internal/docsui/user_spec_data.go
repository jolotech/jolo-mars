package docsui

func UserGroup() Group {
	return Group{
		ID:    "users",
		Title: "Users",
		Sections: []Section{
			{
				ID:    "users-auth",
				Title: "Auth",
				Endpoints: []Endpoint{
					{
						ID: "user-register", 
						Method: "POST",
						Path: "/v1/auth/register",
						Summary: "Registration",
						Auth: "none",
						Usage: &UsageSpec{
							Title: "Usage",
							Notes: []string{
								"This endpoint registers a new user account. The required fields are email or phone, password and otp_option.",
								"Optionally, a guest_id may be provided to associate an existing guest session with the newly created user account. This association ensures session continuity and enables consolidated analytics and activity tracking across both guest and authenticated states.",
								"Upon successful registration, a new user account is created and a verification email or text sms is sent to the provided email address or phone number. The user must verify their email or phone number before they can log in.",
							},
						},
						Request: &RequestSpec{
							ContentType: "application/json",
							Example: map[string]any{
								"name": "jolo delivery",
								"phone": "+2348120618628",
								"otp_option": "email or phone",
								"email": "jolo@example.com",
								"password": "password123",
								"guest_id": "3a9f654f028611f",
							},
						},
						Responses: []ResponseSpec{
							{
								Status: 200,
								Description: "Success",
								Example: map[string]any{
									"status": "success",
									"message": "verification email sent",
									"code": 200,
								},
							},
							{
								Status: 400,
								Description: "Bad Request",
								Example: map[string]any{
									"status": "error",
									"message": "Invalid email or password",
									"code": 400,
								},
							},

						},
					},
					{
						ID: "user-login", 
						Method: "POST", 
						Path: "/v1/auth/login", 
						Summary: "Login", 
						Auth: "none",
						Usage: &UsageSpec{
							Title: "Usage",
							Notes: []string{
								"This endpoint authenticates a user and returns a signed JWT access token upon successful credential validation.",
								"Authentication can be performed using either email or phone in combination with password.",
								"Optionally, a guest_id may be provided to associate an existing guest session with the authenticated user account. This linkage enables session continuity, activity consolidation, and analytics tracking across guest and authenticated states.",
							},
						},
						Request: &RequestSpec{
							ContentType: "application/json",
							Example: map[string]any{
								"email": "jolo@example.com",
								"password": "password123",
								"guest_id": "3a9f654f028611f",
							},
						},
						Responses: []ResponseSpec{
							{
								Status: 200,
								Description: "Login Success", 
								Example: map[string]any{
									"status": "success",
								    "message": "login successful",
									"data": map[string]any{
										"user": map[string]any{
											"public_id ": "3a9f654f028611f",
											"f_name": "jolo",
											"l_name": "user",
											"email": "jolo@gmail.com",
											"phone": "+2348120618620",
											"ref_by": nil,
											"ref_code": "RPC-5NU2-YE87-1O",
											"status": true,
											"is_new": true,
											"is_phone_verified": false,
											"is_email_verified": true,
											"cm_firebase_token": nil,
											"created_at": "2026-01-28T19:58:07.429Z",
											"updated_at": "2026-01-28T22:53:09.321Z",
										},
										"token": "JWT token e.g eyJhbGciOiJIUzI1NiI.......",
									},
									"code": 200,
								},
							},
							{
								Status: 401, 
								Description: "Invalid Credentials",
								Example: map[string]any{
									"status": "error",
									"message": "User credential does not match",
									"code": 401,
								},
							},
						},
					},
					{
						ID: "user-verify-otp", 
						Method: "POST", Path: "/v1/auth/verify-otp", 
						Summary: "Verify OTP",
						Auth: "bearer",
						Usage: &UsageSpec{
							Title: "Usage",
							Notes: []string{
								"This endpoint verifies the OTP sent to the user's email or phone during registration or password reset. The user must provide the OTP code and the associated email or phone number for verification.",
						    },
						},
						Request: &RequestSpec{
								ContentType: "application/json",
								Example: map[string]any{
									"otp": "123456",
									"email": "jolo@example.com",
									"verification_method": "email",
								},
						},
						
						Responses: []ResponseSpec{
							{
								Status: 200,
								Description: "Verification Success",
								Example: map[string]any{
									"status": "success",
									"message": "Verification successful",
									"data": map[string]any{
										"token": "JWT token e.g eyJhbGciOiJIUzI1NiI.......",
										"user": map[string]any{
											"public_id": "AMbm48jjMPaTz0Q",
											"f_name": "jolo",
											"l_name": "delivery",
											"email": "jolodelivery@gmail.com",
											"phone": "+2348120618617",
											"ref_by": nil,
											"ref_code": "RPC-WM5A-FDK8-G3",
											"status": true,
											"is_phone_verified": false,
											"is_email_verified": true,
											"cm_firebase_token": nil,
											"created_at": "2026-01-27T06:16:23.952+01:00",
											"updated_at": "2026-01-27T06:16:43.377+01:00",
										},
									},
									"code": 200,
								},
							},
							{
								Status: 400,
								Description: "Invalid OTP",
								Example: map[string]any{
									"status": "error",
									"message": "Invalid OTP code",
									"code": 400,
								},
							},
						},
					},
					{
						ID: "user-resend-otp",
						Method: "POST",
						Path: "/v1/auth/resend-otp",
						Summary: "Resend OTP",
						Auth: "none",
						Usage: &UsageSpec{
							Title: "Usage",
							Notes: []string{
								"This endpoint triggers the regeneration and delivery of a new One-Time Password (OTP) for verification purposes when the previous OTP has expired or was not delivered. Required parameters include verification_method and either email or phone, based on the specified delivery channel.",
							},
						},
						Request: &RequestSpec{
							ContentType: "application/json",
							Example: map[string]any{
								"verification_method": "email",
								"email or phone":      "jolo@gmail.com",
							},
						},
						Responses: []ResponseSpec{
							{
								Status: 200,
								Description: "OTP Sent",
								Example: map[string]any{
									"status": "success",
									"message": "OTP sent successfully",
									"code": 200,
								},
							},
							{
								Status: 400,
								Description: "Deactivated TOP",
								Example: map[string]any{
									"status": "error",
									"message": "OTP deactivated",
									"error": "Invalid OTP",
									"code": 400,
								},
							},
						},
					},
					{
						ID: "user-forgot-password", 
						Method: "POST", 
						Path: "/v1/auth/forget-password", 
						Summary: "Forgot Password",
						Auth: "none",
						Usage: &UsageSpec{
							Title: "Usage",
							Notes: []string{
								"This endpoint triggers the password recovery workflow. It generates a time-bound OTP and delivers it through the specified verification_method (e.g., email or phone). Required parameters include email and verification_method. After receiving the OTP, the user must call the reset endpoint to validate the OTP and complete the password update process.",
							},
						},
						Request: &RequestSpec{
							ContentType: "application/json",
							Example: map[string]any{
								"email": "jolodelivry@gmail.com",
							    "verification_method": "email",

								"reset_token": "123456",
							    // "email": "jolodelivry@gmail.com",
							    // "verification_method": "email",
							    "password": "1234567890",
							    "confirm_password": "1234567890",
						    },
					    },
						Responses: []ResponseSpec{
							{
								Status: 200,
								Description: "OTP Sent",
								Example: map[string]any{
									"status": "success",
									"message": "OTP sent successfully",
									"code": 200,
								},
							},
						},
					},
				},
			},
			{
				ID:    "users-cart",
				Title: "Cart (pending)",
				Endpoints: []Endpoint{
					{ID: "cart-get", Method: "GET", Path: "/users/cart", Summary: "Get Cart", Auth: "bearer"},
					{ID: "cart-add", Method: "POST", Path: "/users/cart/items", Summary: "Add Item", Auth: "bearer"},
				},
			},
		},
	}
}
