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
						ID: "user-login", 
						Method: "POST", 
						Path: "/v1/auth/login", 
						Summary: "User Login", 
						Auth: "none",
						Usage: &UsageSpec{
							Title: "Usage",
							Notes: []string{
								"Use this endpoint to authenticate a user and get a JWT token.",
								"You can login with either phone or email along with password.",
								"You can also include guest_id field if the user has a guest session to link the session to the user account. also used for analytics and tracking of ther guest and user activities.",
							},
						},
						Request: &RequestSpec{
							ContentType: "application/json",
							Example: map[string]any{
								"email": "jolo@example.com",
								"password": "password123",
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
