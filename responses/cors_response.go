package responses

type CORSResponse struct {
	Headers map[string]string `json:"-"`
	Cookies map[string]string `json:"-"`
}

func (resp *CORSResponse) AddHeader(key, value string) {
	resp.Headers[key] = value
}

func (resp *CORSResponse) AddAllHeaders() {
	if resp.Headers == nil {
		resp.Headers = make(map[string]string)
	}
}
