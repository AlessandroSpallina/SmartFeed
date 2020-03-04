package handler

// RequestBody - received from InfoColumn
type RequestBody struct {
	Username      string `json:"username"`
	ResponseTopic string `json:"response-topic"`
}
