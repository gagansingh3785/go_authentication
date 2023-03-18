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

func AuthoriseSession(next func(http.ResponseWriter, *http.Request, string, string)) requestHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request Payload: %+v\n", r)
		sessionCookie, err := r.Cookie(constants.SESSION_COOKIE)
		if err != nil {
			resp := responses.CommonResponse{
				Error:   constants.InvalidCredentials,
				Message: constants.InvalidCredentials,
			}
			handlers.WriteResponse(w, http.StatusUnauthorized, resp, map[string]string{}, map[string]string{})
			return
		}
		w.Header().Add(constants.SESSION_COOKIE, sessionCookie.Value)
		username, sessionID := parseSessionHeader(sessionCookie.Value)
		user, err := repository.GetUserByUsername(username)
		if err != nil {
			if err == constants.ErrSQLNoRows {
				resp := responses.CommonResponse{
					Error:   constants.InvalidCredentials,
					Message: constants.InvalidCredentials,
				}
				handlers.WriteResponse(w, http.StatusUnauthorized, resp, map[string]string{}, map[string]string{})
			} else {
				resp := responses.CommonResponse{
					Error:   constants.InternalServerError,
					Message: constants.InternalServerError,
				}
				handlers.WriteResponse(w, http.StatusInternalServerError, resp, map[string]string{}, map[string]string{})
			}
		}
		session, err := repository.GetSessionFromUserID(user.UUID)
		switch err {
		case constants.ErrSQLNoRows:
			resp := responses.CommonResponse{
				Error:   constants.SessionExpired,
				Message: constants.SessionExpired,
			}
			handlers.WriteResponse(w, http.StatusUnauthorized, resp, map[string]string{}, map[string]string{})
		case nil:
			if session.SessionID != sessionID {
				resp := responses.CommonResponse{
					Error:   constants.SessionExpired,
					Message: constants.SessionExpired,
				}
				handlers.WriteResponse(w, http.StatusUnauthorized, resp, map[string]string{}, map[string]string{})
				return
			}
			next(w, r, session.SessionID, username)
		default:
			resp := responses.CommonResponse{
				Error:   constants.InternalServerError,
				Message: constants.InternalServerError,
			}
			handlers.WriteResponse(w, http.StatusInternalServerError, resp, map[string]string{}, map[string]string{})
		}
	}
}

func parseSessionHeader(sessionCookieString string) (string, string) {
	username := ""
	sessionID := ""
	sessionCookieDecoded, err := base64.StdEncoding.DecodeString(sessionCookieString)
	fmt.Println("Session Header Decoded: ", sessionCookieDecoded)
	sessionCookie, err := decryptData(sessionCookieDecoded)
	fmt.Println("Session Header Decrypted: ", sessionCookie)
	if err != nil {
		fmt.Println(err.Error())
	}
	parseResult := strings.Split(sessionCookie, ":")
	if len(parseResult) != 2 {
		return username, sessionID
	}
	username, sessionID = parseResult[0], parseResult[1]
	fmt.Println("Username and SessionID: ", username, sessionID)
	return username, sessionID
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
