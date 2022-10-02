package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"os"
)

func main() {
	New()
}

func New() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pdf := gofpdf.New("P", "mm", "A4", pwd+"/font/")
	pdf.AddFont("Helvetica", "", "helvetica_1251.json")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 16)
	tr := pdf.UnicodeTranslatorFromDescriptor("cp1251")
	pdf.Cell(15, 50, tr("Слава Україні"))
	pdf.OutputFileAndClose("test.pdf")
}
