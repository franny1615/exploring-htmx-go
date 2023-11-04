package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

type PageContent struct {
	Title   string
	Content template.HTML
}

func renderPage(pageTitle string, page string, writer http.ResponseWriter, data any) {
	pageTemplate, err := getHtmlTemplateForPage(page, data)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	templatePage, err := template.ParseFiles("html/template.html")

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	templateData := &PageContent{Title: pageTitle, Content: template.HTML(pageTemplate)}
	templatePage.Execute(writer, templateData)
}

func getHtmlTemplateForPage(page string, data any) (template.HTML, error) {
	pageName := fmt.Sprintf("%s.html", page)
	htmlPage := fmt.Sprintf("html/%s", pageName)
	pageTemplate, err := template.ParseFiles(htmlPage)

	if err != nil {
		return template.HTML([]byte("")), err
	}

	pageTemplateCompiled := bytes.Buffer{}
	err = pageTemplate.ExecuteTemplate(&pageTemplateCompiled, pageName, data)

	if err != nil {
		return template.HTML([]byte("")), err
	}

	return template.HTML(pageTemplateCompiled.Bytes()), nil
}
