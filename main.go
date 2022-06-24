package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	pdfGenerator := DefaultPdfGenerator()

	log.Println("Page1")
	// html := `<html><div style="color: #f00;">憑證<h1>時間：2022/01/01</h1></div></html>`
	html := `<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
			<title>合約</title>
		</head>
		<body>
			<h1>這是同意書</h1>
			<p>你好</p>
		</body>
	</html>`
	// page := NewPage(html)
	page := NewPageWithFooter(html)
	pdfGenerator.AddPage(page)
	CreateFile(pdfGenerator, "test123.pdf")

	// TryMerge

	certificateInfo := `
		<!DOCTYPE html>
		<html>
			<head>
				<meta charset="utf-8" />
				<title>憑證</title>
			</head>
			<body>
				<h3>憑證</h3>
				<p>序號：eq24u9r2uqw</p>
				<p>時間：2022/01/01</p>
			</body>
		</html>
	`

	pdfGenerator2 := DefaultPdfGenerator()
	page = NewPageByMerge(html, certificateInfo)
	pdfGenerator2.AddPage(page)
	CreateFile(pdfGenerator2, "merged.pdf")

	log.Println("Done")
}
