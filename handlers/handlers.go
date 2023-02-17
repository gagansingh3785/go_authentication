package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/requests"
	"github.com/gagansingh3785/go_authentication/responses"
	"github.com/gagansingh3785/go_authentication/services"
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
	loginRequest := requests.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		fmt.Println("Error while unmarshalling: ", err.Error())
		handlerResp := responses.LoginResponse{
			Message: "",
			Error:   "Please provide all the required fields",
		}
		WriteResponse(w, http.StatusBadRequest, handlerResp)
		return
	}
	err = loginRequest.Validate()
	if err != nil {
		fmt.Println("Error while unmarshalling: ", err.Error())
		handlerResp := responses.LoginResponse{
			Message: "",
			Error:   "Please provide all the required fields",
		}
		WriteResponse(w, http.StatusBadRequest, handlerResp)
		return
	}
	resp := services.LoginService(loginRequest)
	if resp.Error != constants.EMPTY_STRING {
		WriteResponse(w, http.StatusInternalServerError, resp)
		return
	}
	resp.AddAllHeaders()
	WriteResponse(w, http.StatusOK, resp)
}

func GenerateSessionHandler(w http.ResponseWriter, r *http.Request) {
	generateSessionRequest := requests.GenerateSessionRequest{}
	err := json.NewDecoder(r.Body).Decode(&generateSessionRequest)
	if err != nil {
		handlerResp := responses.LoginResponse{
			Message: "",
			Error:   "Please provide all the required fields",
		}
		WriteResponse(w, http.StatusBadRequest, handlerResp)
		return
	}
	err = generateSessionRequest.Validate()
	if err != nil {
		handlerResp := responses.LoginResponse{
			Message: "",
			Error:   "Please provide all the required fields",
		}
		WriteResponse(w, http.StatusBadRequest, handlerResp)
		return
	}
	resp := services.GenerateSessionService(generateSessionRequest)
	if resp.Error != constants.EMPTY_STRING {
		WriteResponse(w, http.StatusInternalServerError, resp)
	}
	WriteResponse(w, http.StatusCreated, resp)
}

// Handling mail
func SendMail(w http.ResponseWriter, r *http.Request) {
	// Response Declaration
	var handlerResp responses.SendMailResponse

	// Request processing
	sendMailRequest := &requests.SendMailRequest{}
	err := json.NewDecoder(r.Body).Decode(sendMailRequest)

	if err != nil {
		fmt.Println("Error while unmarshalling: ", err.Error())
		handlerResp = responses.SendMailResponse{
			Message: "",
			Error:   "Please provide all the required fields",
		}
		WriteResponse(w, http.StatusBadRequest, handlerResp)
		return
	}

	//logging request
	//fmt.Println(sendMailRequest)

	// Domain processing and making API call to MailGrid
	//url := config.GlobalConfig.SendGrid.APIHost + config.GlobalConfig.SendGrid.APIEndpoint
	//contentType := constants.CONTENT_TYPE
	//receiverName := "gagan"
	//receiverEmail := "9592951585g@gmail.com"

	//GridMailRequest := requests.GetSendGridRequestBody(sendMailRequest.Name,
	//	receiverName,
	//	sendMailRequest.Email,
	//	receiverEmail,
	//	config.GlobalConfig.SendGrid.APIKey,
	//	constants.CONTENT_TYPE,
	//	sendMailRequest.Message,
	//)
	//req, err := json.Marshal(GridMailRequest)
	//bodyReader := bytes.NewReader(req)
	//fmt.Println("url:" + url)
	//response, err := http.Post(url, contentType, bodyReader)
	//if err != nil {
	//	fmt.Println("here2")
	//	fmt.Fprintf(w, err.Error())
	//	return
	//}
	//if response.StatusCode != 200 || response.StatusCode != 201 {
	//	fmt.Println("here3")
	//	fmt.Println(response.Status)
	//	fmt.Fprintf(w, response.Status)
	//	return
	//}
	// Response processing
	handlerResp = responses.SendMailResponse{
		Message: "Success! We will get back to you soon :)",
		Error:   "",
	}
	WriteResponse(w, http.StatusAccepted, handlerResp)
}

func CorsHandler(w http.ResponseWriter, r *http.Request) {
	WriteResponse(w, 200, "")
}

func WriteResponse(w http.ResponseWriter, status int, response any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("%v", w.Header())
}
