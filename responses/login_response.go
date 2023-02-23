package responses

type LoginResponse struct {
	Headers      map[string]string `json:"-"`
	Salt         string            `json:"salt"`
	PasswordHash string            `json:"password_hash"`
	Message      string            `json:"message"`
	Error        string            `json:"error"`
}

func (resp *LoginResponse) AddHeader(key, value string) {
	resp.Headers[key] = value
}

func (resp *LoginResponse) AddCORSHeaders() {
	if resp.Headers == nil {
		resp.Headers = make(map[string]string)
	}
	resp.AddHeader("Access-Control-Allow-Origin", "*")
	resp.AddHeader("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	resp.AddHeader("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func (resp *LoginResponse) AddAllHeaders() {
	if resp.Headers == nil {
		resp.Headers = make(map[string]string)
	}
	resp.AddCORSHeaders()
	resp.AddHeader("Content-Type", "json")
}
