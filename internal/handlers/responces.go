package handlers

// ServerError example
type ServerError struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"status server error"`
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// HTTPSuccess example
type HTTPSuccess struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"OK"`
}
