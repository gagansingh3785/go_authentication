package middleware

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/handlers"
	"github.com/gagansingh3785/go_authentication/repository"
	"github.com/gagansingh3785/go_authentication/responses"
	"net/http"
	"os"
	"strings"
)

type requestHandlerFunc func(http.ResponseWriter, *http.Request)

var randBytes = []byte{32, 12, 45, 54, 67, 42, 23, 200, 101, 234, 12, 222, 39, 91, 87, 45}

func AuthoriseSession(next requestHandlerFunc) requestHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionHeader := r.Header.Get(constants.SESSION_COOKIE)
		if sessionHeader == "" {
			resp := responses.CommonResponse{
				Error:   constants.InvalidCredentials,
				Message: "",
			}
			handlers.WriteResponse(w, http.StatusUnauthorized, resp, map[string]string{})
			return
		}
		username, sessionID := parseSessionHeader(sessionHeader)
		user, err := repository.GetUserByUsername(username)
		if err != nil {
			if err == constants.ErrSQLNoRows {
				resp := responses.CommonResponse{
					Error:   constants.InvalidCredentials,
					Message: "",
				}
				handlers.WriteResponse(w, http.StatusUnauthorized, resp, map[string]string{})
			}
		}
		session, err := repository.GetSessionFromUserID(user.UUID)
		switch err {
		case constants.ErrSQLNoRows:
			resp := responses.CommonResponse{
				Error:   constants.SessionExpired,
				Message: constants.SessionExpired,
			}
			handlers.WriteResponse(w, http.StatusUnauthorized, resp, map[string]string{})
		case nil:
			if session.SessionID != sessionID {
				resp := responses.CommonResponse{
					Error:   constants.SessionExpired,
					Message: constants.SessionExpired,
				}
				handlers.WriteResponse(w, http.StatusUnauthorized, resp, map[string]string{})
			} else {
				next(w, r)
				sessionCookie, err := encryptSessionHeader(w.Header().Get(constants.SESSION_COOKIE))
				if err != nil {
					resp := responses.CommonResponse{
						Error:   constants.InternalServerError,
						Message: constants.InternalServerError,
					}
					handlers.WriteResponse(w, http.StatusInternalServerError, resp, map[string]string{})
					return
				}
				w.Header().Set(constants.SESSION_COOKIE, sessionCookie)
			}
		default:
			resp := responses.CommonResponse{
				Error:   constants.InternalServerError,
				Message: "",
			}
			handlers.WriteResponse(w, http.StatusInternalServerError, resp, map[string]string{})
		}
	}
}

func encryptSessionHeader(sessionCookie string) (string, error) {
	return encryptData(sessionCookie)
}

func parseSessionHeader(sessionHeaderString string) (string, string) {
	username := ""
	sessionID := ""
	sessionHeaderDecoded, err := base64.StdEncoding.DecodeString(sessionHeaderString)
	fmt.Println("Session Header Decoded: ", sessionHeaderDecoded)
	sessionHeader, err := decryptData(sessionHeaderDecoded)
	fmt.Println("Session Header Decrypted: ", sessionHeader)
	if err != nil {
		fmt.Println(err.Error())
	}
	parseResult := strings.Split(sessionHeader, ":")
	if len(parseResult) != 2 {
		return username, sessionID
	}
	username, sessionID = parseResult[0], parseResult[1]
	fmt.Println("Username and SessionID: ", username, sessionID)
	return username, sessionID
}

func encryptData(data string) (string, error) {
	mySecret := os.Getenv(constants.MY_SECRET)
	block, err := aes.NewCipher([]byte(mySecret))
	if err != nil {
		return "", err
	}
	fmt.Println("Session Header before encryption: ", data)
	plainText := []byte(data)
	cfb := cipher.NewCFBEncrypter(block, randBytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	sessionHeaderEncoded := base64.StdEncoding.EncodeToString(cipherText)
	fmt.Println("Session Header after encryption: ", sessionHeaderEncoded)
	return sessionHeaderEncoded, nil
}

func decryptData(data []byte) (string, error) {
	mySecret := os.Getenv(constants.MY_SECRET)
	block, err := aes.NewCipher([]byte(mySecret))
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBDecrypter(block, randBytes)
	plainText := make([]byte, len(data))
	cfb.XORKeyStream(plainText, []byte(data))
	return string(plainText), nil
}
