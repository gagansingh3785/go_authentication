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

type CommonResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func getSessionCookie(username, sessionID string) string {
	sessionHeaderString := username + ":" + sessionID
	fmt.Println(sessionHeaderString)
	sessionHeader, err := encryptSessionCookie(sessionHeaderString)
	fmt.Println(sessionHeader)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return sessionHeader
}

func encryptSessionCookie(sessionCookie string) (string, error) {
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
