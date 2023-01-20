package handlers

import (
	"fmt"
	"github.com/gagansingh3785/go_authentication/constants"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	//This is how templates are executed in golang
	path, err := constants.GetTemplatePath()
	if err != nil {
		fmt.Fprintf(w, "Response:%s Status:%s", constants.InternalServerError, http.StatusInternalServerError)
	}
	t, err := template.New("Home").ParseFiles(path + "/" + constants.HomeHTMLFile)
	if err != nil {
		fmt.Fprintf(w, "Response:%s Status:%s", constants.InternalServerError, http.StatusInternalServerError)
	}
	err = t.Execute(w, "")
	fmt.Println(err)
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

