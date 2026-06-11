package user

type UpdatePasswordRequest struct {
	Email       *string `json:"email" validate:"omitempty,email"`
	CurrentPass *string `json:"current_pass" validate:"omitempty"`
	NewPassword *string `json:"new_password" validate:"omitempty,min=8,max=72"`
}
