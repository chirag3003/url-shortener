package request

// UpdateUserRequest represents the request body for updating user profile.
type UpdateUserRequest struct {
	Name      string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	AvatarURL string `json:"avatarUrl,omitempty" validate:"omitempty,url"`
}
