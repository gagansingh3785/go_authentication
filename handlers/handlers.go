package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/requests"
	"github.com/gagansingh3785/go_authentication/responses"
	"github.com/gagansingh3785/go_authentication/services"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Home page called with requestPayload: %+v\n", r)
	homeRequest := requests.HomeRequest{}
	pageNumber, err := strconv.Atoi(r.URL.Query().Get("pageNumber"))
	if err != nil {
		resp := responses.NewHomeResponse()
		resp.Error = constants.BadRequest
		resp.Message = constants.BadRequest
		WriteResponse(w, http.StatusBadRequest, resp, resp.Headers, resp.Cookies)
		return
	}
	homeRequest.PageNumber = pageNumber
	err = homeRequest.Validate()
	if err != nil {
		resp := responses.NewHomeResponse()
		resp.Error = constants.BadRequest
		resp.Message = constants.BadRequest
		WriteResponse(w, http.StatusBadRequest, resp, resp.Headers, resp.Cookies)
		return
	}
	resp := services.HomeService(homeRequest)
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
		registerResponse := responses.NewRegisterResponse()
		registerResponse.Error = constants.BadRequest
		registerResponse.Message = constants.BadRequest
		WriteResponse(w, http.StatusBadRequest, registerResponse, registerResponse.Headers, registerResponse.Cookies)
	}

	err = registerRequest.Validate()

	if err != nil {
		fmt.Println("Error while unmarshalling: ", err.Error())
		registerResponse := responses.NewRegisterResponse()
		registerResponse.Error = constants.BadRequest
		registerResponse.Message = constants.BadRequest
		WriteResponse(w, http.StatusBadRequest, registerResponse, registerResponse.Headers, registerResponse.Cookies)
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
		loginResponse := responses.NewLoginResponse()
		loginResponse.Error = constants.BadRequest
		loginResponse.Message = constants.BadRequest
		WriteResponse(w, http.StatusBadRequest, loginResponse, loginResponse.Headers, loginResponse.Cookies)
		return
	}
	fmt.Printf("\n %+v \n", loginRequest)
	err = loginRequest.Validate()
	if err != nil {
		fmt.Println("Error while unmarshalling: ", err.Error())
		loginResponse := responses.NewLoginResponse()
		loginResponse.Error = constants.BadRequest
		loginResponse.Message = constants.BadRequest
		WriteResponse(w, http.StatusBadRequest, loginResponse, loginResponse.Headers, loginResponse.Cookies)
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
		generateSessionResponse := responses.NewGenerateSessionResponse()
		generateSessionResponse.Error = constants.BadRequest
		generateSessionResponse.Message = constants.BadRequest
		WriteResponse(w, http.StatusBadRequest, generateSessionResponse,
			generateSessionResponse.Headers,
			generateSessionResponse.Cookies)
		return
	}
	fmt.Printf("\n %+v \n", generateSessionRequest)
	err = generateSessionRequest.Validate()
	if err != nil {
		generateSessionResponse := responses.NewGenerateSessionResponse()
		generateSessionResponse.Error = constants.BadRequest
		generateSessionResponse.Message = constants.BadRequest
		WriteResponse(w, http.StatusBadRequest, generateSessionResponse,
			generateSessionResponse.Headers,
			generateSessionResponse.Cookies)
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
	fmt.Printf("%+v \n", writeRequest)
	if err != nil {
		writeResponse := responses.NewWriteResponse()
		writeResponse.Error = constants.BadRequest
		writeResponse.Message = constants.BadRequest
		WriteResponse(w, http.StatusBadRequest, writeResponse, writeResponse.Headers, writeResponse.Cookies)
		return
	}
	err = writeRequest.Validate()
	if err != nil {
		writeResponse := responses.NewWriteResponse()
		writeResponse.Error = constants.BadRequest
		writeResponse.Message = constants.BadRequest
		WriteResponse(w, http.StatusBadRequest, writeResponse, writeResponse.Headers, writeResponse.Cookies)
		return
	}
	resp := services.WriteService(writeRequest, username)
	addCORSHeaders(resp.Headers)
	switch resp.Error {
	case constants.InternalServerError:
		WriteResponse(w, http.StatusInternalServerError, resp, resp.Headers, resp.Cookies)
	default:
		WriteResponse(w, http.StatusCreated, resp, resp.Headers, resp.Cookies)
	}
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	detailRequest := requests.GetDetailRequest{}
	vars := mux.Vars(r)
	detailRequest.ArticleUUID = vars["articleID"]
	err := detailRequest.Validate()
	if err != nil {
		getDetailResponse := responses.NewGetDetailResponse()
		getDetailResponse.Error = constants.BadRequest
		getDetailResponse.Message = constants.BadRequest
		WriteResponse(w, http.StatusBadRequest, getDetailResponse, getDetailResponse.Headers, getDetailResponse.Cookies)
		return
	}
	resp := services.GetDetailService(detailRequest)
	switch resp.Error {
	case constants.BadRequest:
		WriteResponse(w, http.StatusBadRequest, resp, resp.Headers, resp.Cookies)
	case constants.InternalServerError:
		WriteResponse(w, http.StatusInternalServerError, resp, resp.Headers, resp.Cookies)
	default:
		WriteResponse(w, http.StatusOK, resp, resp.Headers, resp.Cookies)
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
		w.Header().Set(key, value)
	}
	fmt.Printf("%+v \n", headers)
	fmt.Printf("%+v \n", w.Header())
	// This method can only be called once per request
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func addCORSHeaders(headers map[string]string) {
	fmt.Println("here")
	headers["Access-Control-Allow-Origin"] = "http://localhost:3000"
	headers["Access-Control-Allow-Methods"] = "POST, GET, OPTIONS, PUT, DELETE"
	headers["Access-Control-Allow-Headers"] = "Accept, Content-Type, Content-Length, Authorization"
	headers["Access-Control-Expose-Headers"] = "*"
}
