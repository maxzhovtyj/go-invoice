package goinvoice

import "fmt"

func (doc *Document) appendTotal() {
	doc.pdf.SetY(doc.pdf.GetY() + 5)
	doc.pdf.SetFont("Helvetica", "", 12)
	doc.pdf.SetTextColor(35, 35, 35)

	// Draw TOTAL title
	doc.pdf.SetX(120)
	doc.pdf.SetFillColor(212, 212, 212)
	doc.pdf.Rect(120, doc.pdf.GetY(), 40, 10, "F")
	doc.pdf.CellFormat(38,
		10,
		doc.UnicodeTranslatorFunc("Загальна сума"),
		"0",
		0,
		"C",
		false,
		0,
		"",
	)

	// Draw TOTAL value
	doc.pdf.SetX(162)
	doc.pdf.SetFillColor(232, 232, 232)
	doc.pdf.Rect(160, doc.pdf.GetY(), 40, 10, "F")
	doc.pdf.CellFormat(
		40,
		10,
		doc.UnicodeTranslatorFunc(fmt.Sprintf("%f", doc.TotalSum)),
		"0",
		0,
		"C",
		false,
		0,
		"",
	)
}
