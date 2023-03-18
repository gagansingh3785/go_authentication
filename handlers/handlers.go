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

func Home(w http.ResponseWriter, r *http.Request, sessionKey, username string) {
	fmt.Println("Home page called with sessionCookie: ", sessionKey)
	homeRequest := requests.HomeRequest{}
	err := json.NewDecoder(r.Body).Decode(&homeRequest)
	if err != nil {
		resp := responses.HomeResponse{
			Error:   constants.BadRequest,
			Message: constants.BadRequest,
		}
		WriteResponse(w, http.StatusBadRequest, resp, resp.Headers, resp.Cookies)
		return
	}
	err = homeRequest.Validate()
	if err != nil {
		resp := responses.HomeResponse{
			Error:   constants.BadRequest,
			Message: constants.BadRequest,
		}
		WriteResponse(w, http.StatusBadRequest, resp, resp.Headers, resp.Cookies)
		return
	}
	resp := services.HomeService(homeRequest, username, sessionKey)
	switch resp.Error {
	case constants.ArticlePageNotFound:
		WriteResponse(w, http.StatusBadRequest, resp, resp.Headers, resp.Cookies)
	case constants.EMPTY_STRING:
		WriteResponse(w, http.StatusOK, resp, resp.Headers, resp.Cookies)
	default:
		WriteResponse(w, http.StatusInternalServerError, resp, resp.Headers, resp.Cookies)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register route called")
	registerRequest := requests.RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&registerRequest)
	if err != nil {
		fmt.Println("Error while unmarshalling: ", err.Error())
		handlerResp := responses.RegisterResponse{
			Message: "",
			Error:   "Please provide all the required fields",
		}
		WriteResponse(w, http.StatusBadRequest, handlerResp, handlerResp.Headers, handlerResp.Cookies)
	}

	err = registerRequest.Validate()

	if err != nil {
		fmt.Println("Error while unmarshalling: ", err.Error())
		handlerResp := responses.RegisterResponse{
			Message: "",
			Error:   "Please provide all the required fields",
		}
		WriteResponse(w, http.StatusBadRequest, handlerResp, handlerResp.Headers, handlerResp.Cookies)
	}

	resp := services.RegisterService(registerRequest)

	switch resp.Error {
	case constants.UsernameTaken, constants.EmailAlreadyTaken:
		WriteResponse(w, http.StatusBadRequest, resp, resp.Headers, resp.Cookies)
	case constants.InternalServerError:
		WriteResponse(w, http.StatusInternalServerError, resp, resp.Headers, resp.Cookies)
	default:
		WriteResponse(w, http.StatusCreated, resp, resp.Headers, resp.Cookies)
	}
}
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login route called")
	loginRequest := requests.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		fmt.Println("Error while unmarshalling: ", err.Error())
		handlerResp := responses.LoginResponse{
			Message: "",
			Error:   "Please provide all the required fields",
		}
		WriteResponse(w, http.StatusBadRequest, handlerResp, handlerResp.Headers, handlerResp.Cookies)
		return
	}
	fmt.Printf("\n %+v \n", loginRequest)
	err = loginRequest.Validate()
	if err != nil {
		fmt.Println("Error while unmarshalling: ", err.Error())
		handlerResp := responses.LoginResponse{
			Message: "",
			Error:   "Please provide all the required fields",
		}
		WriteResponse(w, http.StatusBadRequest, handlerResp, handlerResp.Headers, handlerResp.Cookies)
		return
	}
	resp := services.LoginService(loginRequest)

	fmt.Printf("\n %+v \n", resp)

	switch resp.Error {
	case constants.InvalidCredentials:
		WriteResponse(w, http.StatusUnauthorized, resp, resp.Headers, resp.Cookies)
	case constants.EMPTY_STRING:
		WriteResponse(w, http.StatusOK, resp, resp.Headers, resp.Cookies)
	default:
		WriteResponse(w, http.StatusInternalServerError, resp, resp.Headers, resp.Cookies)
	}
}

func GenerateSessionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Generate-Session route called")
	generateSessionRequest := requests.GenerateSessionRequest{}
	err := json.NewDecoder(r.Body).Decode(&generateSessionRequest)
	if err != nil {
		handlerResp := responses.LoginResponse{
			Message: constants.BadRequest,
			Error:   constants.BadRequest,
		}
		WriteResponse(w, http.StatusBadRequest, handlerResp, handlerResp.Headers, handlerResp.Cookies)
		return
	}
	fmt.Printf("\n %+v \n", generateSessionRequest)
	err = generateSessionRequest.Validate()
	if err != nil {
		handlerResp := responses.LoginResponse{
			Message: constants.BadRequest,
			Error:   constants.BadRequest,
		}
		WriteResponse(w, http.StatusBadRequest, handlerResp, handlerResp.Headers, handlerResp.Cookies)
		return
	}
	resp := services.GenerateSessionService(generateSessionRequest)

	fmt.Printf("\n %+v \n", resp)

	switch resp.Error {
	case constants.InvalidCredentials:
		WriteResponse(w, http.StatusUnauthorized, resp, resp.Headers, resp.Cookies)
	case constants.EMPTY_STRING:
		WriteResponse(w, http.StatusOK, resp, resp.Headers, resp.Cookies)
	default:
		WriteResponse(w, http.StatusInternalServerError, resp, resp.Headers, resp.Cookies)
	}
}

// Logout Handler
func Logout(w http.ResponseWriter, r *http.Request, sessionKey, username string) {
	fmt.Println("Logout Handler Called")
	fmt.Println("Session Key: ", sessionKey)
	resp := services.LogoutService(sessionKey)
	switch resp.Error {
	case constants.EMPTY_STRING:
		http.SetCookie(w, &http.Cookie{Name: "SessionID", Value: ""})
		WriteResponse(w, http.StatusOK, resp, resp.Headers, resp.Cookies)
	default:
		WriteResponse(w, http.StatusInternalServerError, resp, resp.Headers, resp.Cookies)
	}
}

func Write(w http.ResponseWriter, r *http.Request, sessionKey, username string) {
	fmt.Printf("Write Handler Called username=%s sessionID=%s\n", username, sessionKey)
	writeRequest := requests.WriteRequest{}
	err := json.NewDecoder(r.Body).Decode(&writeRequest)
	if err != nil {
		resp := responses.WriteResponse{
			Error:   constants.BadRequest,
			Message: constants.BadRequest,
		}
		WriteResponse(w, http.StatusBadRequest, resp, resp.Headers, resp.Cookies)
	}
	err = writeRequest.Validate()
	if err != nil {
		resp := responses.WriteResponse{
			Error:   constants.BadRequest,
			Message: constants.BadRequest,
		}
		WriteResponse(w, http.StatusBadRequest, resp, resp.Headers, resp.Cookies)
	}
	resp := services.WriteService(writeRequest, username)
	switch resp.Error {
	case constants.InternalServerError:
		WriteResponse(w, http.StatusInternalServerError, resp, resp.Headers, resp.Cookies)
	default:
		WriteResponse(w, http.StatusCreated, resp, resp.Headers, resp.Cookies)
	}
}

func CorsHandler(w http.ResponseWriter, r *http.Request) {
	corsResponse := responses.CORSResponse{}
	corsResponse.AddAllHeaders()
	WriteResponse(w, 200, corsResponse, corsResponse.Headers, corsResponse.Cookies)
}

func WriteResponse(w http.ResponseWriter, status int, response any, headers, cookies map[string]string) {
	//Setting Response Headers
	addCORSHeaders(headers)
	for key, value := range headers {
		w.Header().Add(key, value)
	}
	w.Header().Add("X-Session-ID", "SOMETHING")
	//Setting Response Cookies
	for key, value := range cookies {
		cookie := &http.Cookie{
			Name:  key,
			Value: value,
		}
		w.Header().Add(cookie.Name, cookie.Value)
	}
	// This method can only be called once per request
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func addCORSHeaders(headers map[string]string) {
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Access-Control-Allow-Origin"] = "*"
	headers["Access-Control-Allow-Methods"] = "POST, GET, OPTIONS, PUT, DELETE"
	headers["Access-Control-Allow-Headers"] = "Accept, Content-Type, Content-Length, Authorization"
	headers["Access-Control-Expose-Headers"] = "*"
}
