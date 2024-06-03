package utils

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
type UploadResponse struct {
	Message string `json:"message,omitempty"`
	URL     string `json:"url,omitempty"`
}
