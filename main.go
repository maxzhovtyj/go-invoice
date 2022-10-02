package go_invoice

import (
	"github.com/jung-kurt/gofpdf"
	"log"
)

func New() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddFont("Helvetica", "", "helvetica_1251.json")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 16)
	tr := pdf.UnicodeTranslatorFromDescriptor("cp1251")
	pdf.Cell(15, 50, tr("русский текст"))
	err := pdf.OutputFileAndClose("test.pdf")
	if err != nil {
		log.Println(err)
	}
}
