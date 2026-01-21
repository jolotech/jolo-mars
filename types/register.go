package types


type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
	RefCode  string `json:"ref_code"`
}