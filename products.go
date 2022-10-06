package goinvoice

import "fmt"

func (doc *Document) writeProducts() {
	doc.pdf.SetFont(doc.Font, "", 10)

	doc.pdf.SetX(10)
	doc.pdf.SetY(doc.pdf.GetY() + 10)

	for _, p := range doc.Products {
		// Append to pdf
		p.appendColTo(doc)

		if doc.pdf.GetY() > 255 {
			// Add page
			doc.pdf.AddPage()
			doc.drawsTableTitles()
			doc.pdf.SetY(doc.pdf.GetY() + 5)
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
		90,
		3,
		doc.UnicodeTranslatorFunc(p.Title),
		"0",
		"",
		false,
	)

	// Compute line height
	colHeight := doc.pdf.GetY() - baseY

	if p.Packaging == "" {
		p.Packaging = "---"
	}
	// Product packaging
	doc.pdf.SetY(baseY)
	doc.pdf.SetX(100)
	doc.pdf.CellFormat(
		20,
		colHeight,
		doc.UnicodeTranslatorFunc(p.Packaging),
		"0",
		0,
		"C",
		false,
		0,
		"",
	)

	// Unit price
	doc.pdf.SetY(baseY)
	doc.pdf.SetX(120)
	doc.pdf.CellFormat(
		20,
		colHeight,
		doc.UnicodeTranslatorFunc(fmt.Sprintf("%f", p.Price)),
		"0",
		0,
		"C",
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
		"C",
		false,
		0,
		"",
	)

	// Discount
	var discountStr string
	if p.Discount == 0 {
		discountStr = "---"
	} else {
		discountStr = fmt.Sprintf("%f", p.Discount)
	}
	doc.pdf.SetX(160)
	doc.pdf.CellFormat(
		20,
		colHeight,
		doc.UnicodeTranslatorFunc(discountStr),
		"0",
		0,
		"C",
		false,
		0,
		"",
	)

	// total price
	var totalPrice float64
	if p.Total == 0 {
		totalPrice = p.Price * p.Quantity
	} else {
		totalPrice = p.Total
	}

	doc.pdf.SetX(180)
	doc.pdf.CellFormat(
		20,
		colHeight,
		doc.UnicodeTranslatorFunc(fmt.Sprintf("%f", totalPrice)),
		"0",
		0,
		"C",
		false,
		0,
		"",
	)

	// Set Y for next line
	doc.pdf.SetY(baseY + colHeight)
}
