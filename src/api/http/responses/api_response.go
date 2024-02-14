package responses

type ApiResponse struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
