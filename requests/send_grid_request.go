package requests

type MailGridRequest struct {
	//headers
	Authorization string `json:"-"`
	ContentType   string `json:"-"`

	//body
	Personalizations []Personalizations `json:"personalizations"`
	Content          []Content          `json:"Content"`
	From             From               `json:"from"`
}

type Personalizations struct {
	To      []To   `json:"to"`
	Subject string `json:"subject"`
}

type To struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type From struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Content struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func GetSendGridRequestBody(senderName, receiverName, senderEmail, receiverEmail, authorizationToken, contentType, message string) MailGridRequest {
	request := MailGridRequest{}
	request.Authorization = authorizationToken
	request.ContentType = contentType
	personalizations := Personalizations{}
	personalizations.To = []To{
		{
			Email: receiverEmail,
			Name:  receiverName,
		},
	}
	request.Personalizations = []Personalizations{personalizations}
	content := []Content{
		{
			Type:  "text/plain",
			Value: message,
		},
	}
	request.Content = content
	request.From = From{
		Email: senderEmail,
		Name:  senderName,
	}
	return request
}
