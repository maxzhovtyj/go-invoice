package goinvoice

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"time"
)

// TODO add creation date
// TODO add order id
// TODO Think about album orientation

func (doc *Document) BuildPdf() (gofpdf.Fpdf, error) {
	err := doc.validate()
	if err != nil {
		return gofpdf.Fpdf{}, err
	}

	doc.pdf.SetXY(10, 10)

	// Draw text
	doc.pdf.SetFont("Helvetica", "", 16)
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
	return fmt.Sprintf("Замовлення №%s", doc.OrderId)
}

func (doc *Document) writeDivider() {
	doc.pdf.SetX(10)
	doc.pdf.MultiCell(190, 0, doc.UnicodeTranslatorFunc(""), "B", "L", false)
}

func (doc *Document) writeDate() {
	doc.pdf.SetXY(10, doc.pdf.GetY()+10)
	doc.pdf.SetFont("Helvetica", "", 10)
	doc.pdf.MultiCell(190, 5, doc.UnicodeTranslatorFunc(doc.getCurrentTimeString()), "B", "L", false)
}

func (doc *Document) getCurrentTimeString() string {
	return fmt.Sprintf(
		"Видано: %d/%d/%d, %d:%d:%d",
		time.Now().Day(),
		int(time.Now().Month()),
		time.Now().Year(),
		time.Now().Hour(),
		time.Now().Minute(),
		time.Now().Second(),
	)
}
