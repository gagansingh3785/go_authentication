package responses

type CORSResponse struct {
	Headers map[string]string `json:"-"`
}

func (resp *CORSResponse) AddCORSHeaders() {
	resp.AddHeader("Access-Control-Allow-Origin", "*")
	resp.AddHeader("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	resp.AddHeader("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func (resp *CORSResponse) AddHeader(key, value string) {
	resp.Headers[key] = value
}

func (resp *CORSResponse) AddAllHeaders() {
	if resp.Headers == nil {
		resp.Headers = make(map[string]string)
	}
	resp.AddCORSHeaders()
}
