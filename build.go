package goinvoice

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"time"
)

// TODO Think about album orientation
// TODO font option
// TODO language switching

func (doc *Document) BuildPdf() (gofpdf.Fpdf, error) {
	if doc.Font == "" {
		doc.Font = "Arial"
	}
	if doc.Language == "UK" {
		doc.Font = "Helvetica"
	}

	err := doc.validate()
	if err != nil {
		return gofpdf.Fpdf{}, err
	}

	doc.pdf.SetXY(10, 10)

	// Draw text
	doc.pdf.SetFont(doc.Font, "", 16)
	doc.pdf.CellFormat(180, 5, doc.UnicodeTranslatorFunc(doc.orderTitleToString()), "0", 0, "L", false, 0, "")

	if doc.Company != nil {
		doc.writeCompany()
	}

	if doc.Customer != nil {
		doc.writeCustomer()
	}

	doc.writeDate()

	doc.drawsTableTitles()

	if doc.Products != nil {
		doc.writeProducts()
	}

	doc.writeDivider()

	posY := doc.supplierAndReceiver()
	doc.appendTotal(posY)
	return *doc.pdf, nil
}

func (doc *Document) orderTitleToString() string {
	switch doc.Language {
	case "UK":
		return fmt.Sprintf("Замовлення №%s", doc.OrderId)
	default:
		return fmt.Sprintf("Order No.%s", doc.OrderId)
	}
}

func (doc *Document) writeDivider() {
	doc.pdf.SetX(10)
	doc.pdf.MultiCell(190, 0, doc.UnicodeTranslatorFunc(""), "B", "L", false)
}

func (doc *Document) writeDate() {
	doc.pdf.SetFont(doc.Font, "", 10)
	doc.pdf.SetXY(10, doc.pdf.GetY()+10)
	doc.pdf.MultiCell(190, 5, doc.UnicodeTranslatorFunc(doc.getCurrentTimeString()), "B", "L", false)
}

func (doc *Document) getCurrentTimeString() string {
	var issuedStr string
	switch doc.Language {
	case "UK":
		issuedStr = "Видано"
	default:
		issuedStr = "Issued"
	}

	return fmt.Sprintf(
		"%s: %d/%d/%d, %d:%d:%d",
		issuedStr,
		time.Now().Day(),
		int(time.Now().Month()),
		time.Now().Year(),
		time.Now().Hour(),
		time.Now().Minute(),
		time.Now().Second(),
	)
}
