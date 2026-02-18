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
					{ID: "user-login", Method: "POST", Path: "/users/auth/login", Summary: "Login", Auth: "none"},
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
