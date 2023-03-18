package responses

type CORSResponse struct {
	Headers map[string]string `json:"-"`
	Cookies map[string]string `json:"-"`
}

func (resp *CORSResponse) AddHeader(key, value string) {
	resp.Headers[key] = value
}

func (resp *CORSResponse) AddAllHeaders() {

}

func NewCORSResponse() CORSResponse {
	resp := CORSResponse{}
	resp.Headers = make(map[string]string)
	resp.Cookies = make(map[string]string)
	return resp
}