package goinvoice

import "fmt"

func (doc *Document) writeProducts() {
	doc.pdf.SetX(10)
	doc.pdf.SetY(doc.pdf.GetY() + 10)
	doc.pdf.SetFont("Helvetica", "", 10)

	for _, p := range doc.Products {
		// Append to pdf
		p.appendColTo(doc)

		if doc.pdf.GetY() > 260 {
			// Add page
			doc.pdf.AddPage()
			doc.drawsTableTitles()
			doc.pdf.SetY(doc.pdf.GetY() + 5)
			doc.pdf.SetFont("Helvetica", "", 10)
		}

		doc.pdf.SetX(10)
		doc.pdf.SetY(doc.pdf.GetY() + 5)
	}
}

// appendColTo document doc
func (p *Product) appendColTo(doc *Document) {
	// Get base Y (top of line)
	baseY := doc.pdf.GetY()

	// Name
	doc.pdf.SetX(10)
	doc.pdf.MultiCell(
		100,
		3,
		doc.UnicodeTranslatorFunc(p.Title),
		"",
		"",
		false,
	)

	// Compute line height
	colHeight := doc.pdf.GetY() - baseY

	// Unit price
	doc.pdf.SetY(baseY)
	doc.pdf.SetX(120)
	doc.pdf.CellFormat(
		20,
		colHeight,
		doc.UnicodeTranslatorFunc(fmt.Sprintf("%f", p.Price)),
		"0",
		0,
		"",
		false,
		0,
		"",
	)

	// Quantity
	doc.pdf.SetX(140)
	doc.pdf.CellFormat(
		20,
		colHeight,
		doc.UnicodeTranslatorFunc(fmt.Sprintf("%f", p.Quantity)),
		"0",
		0,
		"",
		false,
		0,
		"",
	)

	// Discount
	doc.pdf.SetX(160)
	doc.pdf.CellFormat(
		20,
		colHeight,
		doc.UnicodeTranslatorFunc("---"),
		"0",
		0,
		"",
		false,
		0,
		"",
	)

	// TOTAL TTC
	doc.pdf.SetX(180)
	doc.pdf.CellFormat(
		20,
		colHeight,
		doc.UnicodeTranslatorFunc(fmt.Sprintf("%f", p.Total)),
		"0",
		0,
		"",
		false,
		0,
		"",
	)

	// Set Y for next line
	doc.pdf.SetY(baseY + colHeight)
}
