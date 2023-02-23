package responses

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/gagansingh3785/go_authentication/constants"
	"os"
)

var randBytes = []byte{32, 12, 45, 54, 67, 42, 23, 200, 101, 234, 12, 222, 39, 91, 87, 45}

type GenerateSessionResponse struct {
	Headers map[string]string `json:"-"`
	Message string            `json:"message"`
	Error   string            `json:"error"`
}

func (resp *GenerateSessionResponse) AddHeader(key, value string) {
	resp.Headers[key] = value
}

func (resp *GenerateSessionResponse) AddCORSHeaders() {
	if resp.Headers == nil {
		resp.Headers = make(map[string]string)
	}
	resp.AddHeader("Access-Control-Allow-Origin", "*")
	resp.AddHeader("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	resp.AddHeader("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func (resp *GenerateSessionResponse) AddAllHeaders(username, sessionID string) {
	fmt.Println("Here ^^^^^^^^^^^")
	if resp.Headers == nil {
		resp.Headers = make(map[string]string)
	}
	resp.AddCORSHeaders()
	resp.AddHeader(constants.CONTENT_TYPE, "json")
	sessionHeader := getSessionHeader(username, sessionID)
	resp.AddHeader(constants.SESSION_COOKIE, sessionHeader)
}

func getSessionHeader(username, sessionID string) string {
	sessionHeaderString := username + ":" + sessionID
	fmt.Println(sessionHeaderString)
	sessionHeader, err := encryptSessionHeader(sessionHeaderString)
	fmt.Println(sessionHeader)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return sessionHeader
}

func encryptSessionHeader(sessionCookie string) (string, error) {
	return encryptData(sessionCookie)
}

func encryptData(data string) (string, error) {
	mySecret := os.Getenv(constants.MY_SECRET)
	block, err := aes.NewCipher([]byte(mySecret))
	if err != nil {
		return "", err
	}
	plainText := []byte(data)
	cfb := cipher.NewCFBEncrypter(block, randBytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	sessionHeaderEncoded := base64.StdEncoding.EncodeToString(cipherText)
	return sessionHeaderEncoded, nil
}
