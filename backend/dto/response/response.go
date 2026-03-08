package response

// UserResponse represents the user data returned in API responses.
type UserResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	PhoneNo string `json:"phoneNo,omitempty"`
}

// LoginResponse represents the response after successful login.
type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

// MessageResponse represents a simple message response.
type MessageResponse struct {
	Message string `json:"message"`
}

// MediaUploadResponse represents the response after a successful file upload.
type MediaUploadResponse struct {
	Message string   `json:"message"`
	Files   []string `json:"files"`
}
