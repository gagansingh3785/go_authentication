package constants

import (
	"fmt"
	"os"
)

const (
	TemplateDir       = "/templates"
	HomeHTMLFile      = "home.html"
	MigrationFilePath = "./migrations"
)

func GetTemplatePath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path := cwd + TemplateDir
	fmt.Println(path)
	return path, nil
}
