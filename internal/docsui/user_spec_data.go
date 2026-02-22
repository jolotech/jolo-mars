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
					{ID: "user-forgot-password", Method: "POST", Path: "/users/auth/forgot-password", Summary: "Forgot Password", Auth: "none"},
				},
			},
			{
				ID:    "users-cart",
				Title: "Cart",
				Endpoints: []Endpoint{
					{ID: "cart-get", Method: "GET", Path: "/users/cart", Summary: "Get Cart", Auth: "bearer"},
					{ID: "cart-add", Method: "POST", Path: "/users/cart/items", Summary: "Add Item", Auth: "bearer"},
				},
			},
		},
	}
}
