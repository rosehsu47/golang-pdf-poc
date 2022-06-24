package main

import (
	"bytes"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	log "github.com/sirupsen/logrus"
)

func DefaultPdfGenerator() *wkhtmltopdf.PDFGenerator {

	pdfGenerator, err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		log.Warn("NewPDFGenerator: ", err)
	}

	pdfGenerator.Dpi.Set(600)
	pdfGenerator.NoCollate.Set(false)
	pdfGenerator.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfGenerator.MarginTop.Set(100)

	return pdfGenerator
}

func NewPage(html string) wkhtmltopdf.PageProvider {
	page := wkhtmltopdf.NewPageReader(strings.NewReader(html))

	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(8)

	return page
}

func NewPageWithFooter(html string) wkhtmltopdf.PageProvider {
	page := wkhtmltopdf.NewPageReader(strings.NewReader(html))

	page.FooterHTML.Set("footer.html")
	page.FooterLine.Set(true)
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(8)
	log.Println(page)

	return page
}

func NewPageByMerge(html string, certificateInfo string) *wkhtmltopdf.PageReader {
	log.Println("Merge Html")

	buf := new(bytes.Buffer)
	buf.WriteString(certificateInfo)
	// buf.WriteString(`<P style="page-break-before: always">`)
	buf.WriteString(`<br/>`)
	buf.WriteString(html)

	result := wkhtmltopdf.NewPageReader(buf)
	return result
}

func CreateFile(pdfGenerator *wkhtmltopdf.PDFGenerator, fileName string) error {
	err := pdfGenerator.Create()
	if err != nil {
		log.Println("Create: ", err)
		return err
	}

	err = pdfGenerator.WriteFile(fileName)
	if err != nil {
		log.Println("WriteFile: ", err)
		return err
	}

	log.Println("Success to create file", fileName)
	return nil
}
