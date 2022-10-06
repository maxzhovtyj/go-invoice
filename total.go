package goinvoice

import "fmt"

func (doc *Document) appendTotal(posY float64) {
	var totalTitle string

	switch doc.Language {
	case "UK":
		totalTitle = "Загальна сума"
	default:
		totalTitle = "Total"
	}

	var totalSum float64
	if doc.TotalSum == 0 {
		for _, p := range doc.Products {
			totalSum += p.Price * p.Quantity
		}
	} else {
		totalSum = doc.TotalSum
	}

	doc.pdf.SetFont(doc.Font, "", 12)

	doc.pdf.SetY(posY + 5)
	doc.pdf.SetTextColor(35, 35, 35)

	// Draw TOTAL title
	doc.pdf.SetX(120)
	doc.pdf.SetFillColor(212, 212, 212)
	doc.pdf.Rect(120, doc.pdf.GetY(), 40, 10, "F")
	doc.pdf.CellFormat(38,
		10,
		doc.UnicodeTranslatorFunc(totalTitle),
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
		doc.UnicodeTranslatorFunc(fmt.Sprintf("%f", totalSum)),
		"0",
		0,
		"C",
		false,
		0,
		"",
	)
}
