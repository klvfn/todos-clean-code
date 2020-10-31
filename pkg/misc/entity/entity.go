package entity

// Response standard API response
type Response struct {
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// ErrorResponse error struct when error exist
type ErrorResponse struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	Message     string `json:"message"`
}
