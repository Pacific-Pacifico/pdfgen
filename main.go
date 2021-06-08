package main

import "pdfgenapp/pdfgenerator"

func main() {
	pdfgenerator.GenerateInvoice()
	pdfgenerator.GenerateCertificate()
	// pdfgenerator.getDateInWords()
	// pdfgenerator.getDateTimeInWords()
}
