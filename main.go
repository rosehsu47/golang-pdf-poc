package main

import (
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	pdfGenerator, err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		log.Warn("NewPDFGenerator: ", err)
	}

	pdfGenerator.Dpi.Set(600)
	pdfGenerator.NoCollate.Set(false)
	pdfGenerator.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfGenerator.MarginBottom.Set(40)

	html := "<html>Hi</html>"
	pdfGenerator.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(html)))

	err = pdfGenerator.Create()
	if err != nil {
		log.Println("Create: ", err)
	}

	fileName := "test123.pdf"
	err = pdfGenerator.WriteFile(fileName)
	if err != nil {
		log.Println("WriteFile: ", err)
	}

	log.Println("Done")
}
