package models

// ErrorResponse represents a generic error response
// swagger:model
type ErrorResponse struct {
	// Error message
	Error string `json:"error"`
}
