package response

// APIResponse is the standardized response envelope for all API responses.
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorBody  `json:"error,omitempty"`
}

// ErrorBody contains error details in the API response.
type ErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// OK creates a successful API response with data.
func OK(data interface{}) *APIResponse {
	return &APIResponse{
		Success: true,
		Data:    data,
	}
}

// Err creates an error API response.
func Err(code string, message string) *APIResponse {
	return &APIResponse{
		Success: false,
		Error: &ErrorBody{
			Code:    code,
			Message: message,
		},
	}
}
