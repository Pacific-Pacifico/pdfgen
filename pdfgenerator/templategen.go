package pdfgenerator

import (
	"bytes"
	"log"
	"text/template"
)

func (Idata InvoiceData) RenderTemplate(htmlFile string) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	parsedTemplate, _ := template.ParseFiles(htmlFile)
	err := parsedTemplate.Execute(buf, Idata)
	if err != nil {
		log.Println("Error executing template :", err)
	}
	//fmt.Println(buf)
	return buf, err
}

func (CData CertificateData) RenderTemplate(htmlFile string) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	parsedTemplate, _ := template.ParseFiles(htmlFile)
	err := parsedTemplate.Execute(buf, CData)
	if err != nil {
		log.Println("Error executing template :", err)
	}
	//fmt.Println(buf)
	return buf, err
}

/**
func main() {
	buf := renderTemplate()
	fmt.Println(buf)
}
**/
