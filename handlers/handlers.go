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
	registerRequest := requests.RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&registerRequest)
	if err != nil {
		fmt.Println("Error while unmarshalling: ", err.Error())
		handlerResp := responses.RegisterResponse{
			Message: "",
			Error:   "Please provide all the required fields",
		}
		WriteResponse(w, http.StatusBadRequest, handlerResp, handlerResp.Headers)
	}
	err = registerRequest.Validate()
	if err != nil {
		fmt.Println("Error while unmarshalling: ", err.Error())
		handlerResp := responses.RegisterResponse{
			Message: "",
			Error:   "Please provide all the required fields",
		}
		WriteResponse(w, http.StatusBadRequest, handlerResp, handlerResp.Headers)
	}
	resp := services.RegisterService(registerRequest)
	switch resp.Error {
	case constants.InternalServerError:
		WriteResponse(w, http.StatusInternalServerError, resp, resp.Headers)
	default:
		WriteResponse(w, http.StatusCreated, resp, resp.Headers)
	}
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
		WriteResponse(w, http.StatusBadRequest, handlerResp, handlerResp.Headers)
		return
	}
	err = loginRequest.Validate()
	if err != nil {
		fmt.Println("Error while unmarshalling: ", err.Error())
		handlerResp := responses.LoginResponse{
			Message: "",
			Error:   "Please provide all the required fields",
		}
		WriteResponse(w, http.StatusBadRequest, handlerResp, handlerResp.Headers)
		return
	}
	resp := services.LoginService(loginRequest)

	switch resp.Error {
	case constants.InvalidCredentials:
		WriteResponse(w, http.StatusUnauthorized, resp, resp.Headers)
	case constants.EMPTY_STRING:
		WriteResponse(w, http.StatusOK, resp, resp.Headers)
	default:
		WriteResponse(w, http.StatusInternalServerError, resp, resp.Headers)
	}
}

func GenerateSessionHandler(w http.ResponseWriter, r *http.Request) {
	generateSessionRequest := requests.GenerateSessionRequest{}
	err := json.NewDecoder(r.Body).Decode(&generateSessionRequest)
	if err != nil {
		handlerResp := responses.LoginResponse{
			Message: "",
			Error:   "Please provide all the required fields",
		}
		WriteResponse(w, http.StatusBadRequest, handlerResp, handlerResp.Headers)
		return
	}
	err = generateSessionRequest.Validate()
	if err != nil {
		handlerResp := responses.LoginResponse{
			Message: "",
			Error:   "Please provide all the required fields",
		}
		WriteResponse(w, http.StatusBadRequest, handlerResp, handlerResp.Headers)
		return
	}
	resp := services.GenerateSessionService(generateSessionRequest)

	switch resp.Error {
	case constants.InvalidCredentials:
		WriteResponse(w, http.StatusUnauthorized, resp, resp.Headers)
	case constants.EMPTY_STRING:
		WriteResponse(w, http.StatusOK, resp, resp.Headers)
	default:
		WriteResponse(w, http.StatusInternalServerError, resp, resp.Headers)
	}
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
		WriteResponse(w, http.StatusBadRequest, handlerResp, handlerResp.Headers)
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
	WriteResponse(w, http.StatusAccepted, handlerResp, handlerResp.Headers)
}

func CorsHandler(w http.ResponseWriter, r *http.Request) {
	corsResponse := responses.CORSResponse{}
	corsResponse.AddAllHeaders()
	WriteResponse(w, 200, corsResponse, corsResponse.Headers)
}

func WriteResponse(w http.ResponseWriter, status int, response any, headers map[string]string) {
	for key, value := range headers {
		w.Header().Add(key, value)
	}
	// This method can only be called once per request
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
