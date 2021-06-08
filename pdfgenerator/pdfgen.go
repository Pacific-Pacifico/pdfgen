package pdfgenerator

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type Document interface {
	RenderTemplate(string) (*bytes.Buffer, error)
}

type InvoiceData struct {
	InvoiceNumber   uint64
	Date            string
	CustomerName    string
	CustomerEmail   string
	CustomerPhone   uint64
	Course          string
	OriginalPrice   float32
	DiscountPercent float32
	Discount        float32
	TaxAmount       float32
	TotalAmount     float32
	GrandTotal      float32
	PaymentMode     string
	Sign            string
}

type CertificateData struct {
	CssFile  string
	JsFile   string
	FontFile string

	Username string
	Course   string
	Date     string
	Image    string
	Sign     string
}

func GeneratePDF(docObj Document, htmlFile string, pdfFile string, pdfWidth uint, pdfHeight uint) {
	pg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		fmt.Println("error")
		return
	}

	//data, err := ioutil.ReadFile("template/index.html")
	data, err := docObj.RenderTemplate(htmlFile)

	if err != nil {
		panic(err)
	}

	//htmlStr := string(data)
	htmlStr := data.String()
	// fmt.Print(htmlStr)

	page := wkhtml.NewPageReader(strings.NewReader(htmlStr))
	fmt.Println("***template parsed sucessfully***")
	pg.AddPage(page)
	// pg.PageSize.Set(wkhtml.PageSizeCustom)
	// pg.PageWidth.Set(210)
	// pg.PageHeight.Set(156)
	pg.PageWidth.Set(pdfWidth) //in millimeters
	pg.PageHeight.Set(pdfHeight)
	abspath, err := filepath.Abs("./pdfgenerator/template")
	fmt.Println("path-------" + abspath)
	if err != nil {
		panic(err)
	}
	page.Allow.Set(abspath)
	// page.Allow.Set("D:/pdfgenerator/template")
	fmt.Println(abspath)
	//page.Zoom.Set(0.95)
	// Create PDF document in internal buffer
	err = pg.Create()
	if err != nil {
		log.Fatal(err)
	}

	//Your Pdf Name
	err = pg.WriteFile(pdfFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}

func GenerateInvoice() {
	var IData InvoiceData
	IData.InvoiceNumber = 1234567890 // Transaction number
	IData.Date = GetDateTimeInWords()
	IData.CustomerEmail = "email@gmail.com"
	IData.CustomerName = "abc xyz"
	IData.CustomerPhone = 9876543210
	IData.Course = "web development with html ,css & Javascript"
	IData.OriginalPrice = 299
	IData.TaxAmount = 0
	IData.TotalAmount = 199 - IData.TaxAmount
	IData.Discount = IData.OriginalPrice - IData.TotalAmount
	IData.DiscountPercent = (IData.Discount / IData.OriginalPrice) * 100
	IData.GrandTotal = IData.TotalAmount
	IData.PaymentMode = "UPI"

	s, _ := filepath.Abs("pdfgenerator/template/sign.jpg")
	IData.Sign = s
	GeneratePDF(IData, "pdfgenerator/template/Invoice_sample.html", "invoice.pdf", 210, 290)
}

func GenerateCertificate() {
	var CData CertificateData
	cs, _ := filepath.Abs("pdfgenerator/template/bootstrap.min.css")
	js, _ := filepath.Abs("pdfgenerator/template/bootstrap.min.js")
	img, _ := filepath.Abs("pdfgenerator/template/apni_shiksha_logo.jpg")
	sign, _ := filepath.Abs("pdfgenerator/template/sign.jpg")
	ff, _ := filepath.Abs("pdfgenerator/template/font_file.ttf")
	fmt.Println(cs)
	fmt.Println(js)
	fmt.Println(img)
	fmt.Println(ff)

	CData.CssFile = cs
	CData.JsFile = js
	CData.FontFile = ff

	CData.Image = img
	CData.Sign = sign
	CData.Course = "web development with html ,css & Javascript"
	CData.Username = "Prashant Kumar"
	CData.Date = GetDateInWords()
	GeneratePDF(CData, "pdfgenerator/template/Certificate_sample.html", "certificate.pdf", 210, 156)
}

// func main() {
// 	GenerateInvoice()
// 	GenerateCertificate()
// 	// getDateInWords()
// 	// getDateTimeInWords()
// }
