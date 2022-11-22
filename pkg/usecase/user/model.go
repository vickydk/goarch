package user

type RegisterRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UpdateNameReq struct {
	Name string `json:"name" validate:"required"`
}

type UpdatePasswordReq struct {
	OldPassword string `json:"old_password" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type ResetPasswordReq struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password" validate:"required"`
}
