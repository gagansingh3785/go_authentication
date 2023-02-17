package responses

type GenerateSessionResponse struct {
	Headers map[string]string `json:"-"`
	Message string            `json:"message"`
	Error   string            `json:"error"`
}

func (resp *GenerateSessionResponse) AddHeader(key, value string) {
	resp.Headers[key] = value
}

func (resp *GenerateSessionResponse) AddCORSHeaders() {
	resp.AddHeader("Access-Control-Allow-Origin", "*")
	resp.AddHeader("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	resp.AddHeader("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func (resp *GenerateSessionResponse) AddAllHeaders(sessionID string) {
	resp.AddCORSHeaders()
	resp.AddHeader("Content-Type", "json")
	resp.AddHeader("Cookie", sessionID)
}
