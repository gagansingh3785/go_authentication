package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gagansingh3785/go_authentication/config"
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/requests"
	"github.com/gagansingh3785/go_authentication/responses"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	//This is how templates are executed in golang
	fmt.Fprintf(w, "This is the home page")
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the register page")
}

func Login(w http.ResponseWriter, r *http.Request) {

}

func HandleLogin(w http.ResponseWriter, r *http.Request) {

}

func HandleRegister(w http.ResponseWriter, r *http.Request) {

}

// Handling mail
func SendMail(w http.ResponseWriter, r *http.Request) {
	// Request processing

	sendMailRequest := &requests.SendMailRequest{}
	err := json.NewDecoder(r.Body).Decode(sendMailRequest)
	if err != nil {
		fmt.Println("here1")
		fmt.Fprintf(w, "error")
		return
	}

	// Domain processing and making API call to MailGrid
	url := config.GlobalConfig.SendGrid.APIHost + config.GlobalConfig.SendGrid.APIEndpoint
	contentType := constants.CONTENT_TYPE
	receiverName := "gagan"
	receiverEmail := "9592951585g@gmail.com"

	GridMailRequest := requests.GetSendGridRequestBody(sendMailRequest.Name,
		receiverName,
		sendMailRequest.Email,
		receiverEmail,
		config.GlobalConfig.SendGrid.APIKey,
		constants.CONTENT_TYPE,
		sendMailRequest.Message,
	)
	req, err := json.Marshal(GridMailRequest)
	bodyReader := bytes.NewReader(req)
	fmt.Println("url:" + url)
	response, err := http.Post(url, contentType, bodyReader)
	if err != nil {
		fmt.Println("here2")
		fmt.Fprintf(w, err.Error())
		return
	}
	if response.StatusCode != 200 || response.StatusCode != 201 {
		fmt.Println("here3")
		fmt.Println(response.Status)
		fmt.Fprintf(w, response.Status)
		return
	}
	// Response processing
	handlerResp := responses.SendMailResponse{
		Message: "Success! We will get back to you soon :)",
		Error:   nil,
	}
	WriteResponse(w, http.StatusAccepted, handlerResp)
}

func WriteResponse(w http.ResponseWriter, status int, response any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
