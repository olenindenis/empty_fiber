package handlers

// ServerError example
type ServerError struct {
	Message string `json:"message" example:"status server error"`
}

// HTTPError example
type HTTPError struct {
	Message string `json:"message" example:"status bad request"`
}

// HTTPSuccess example
type HTTPSuccess struct {
	Message string `json:"message" example:"OK"`
}
