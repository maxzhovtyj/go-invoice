package goinvoice

import "fmt"

func (doc *Document) writeProducts() {
	doc.pdf.SetFont(doc.Font, "", 10)

	doc.pdf.SetX(10)
	doc.pdf.SetY(doc.pdf.GetY() + 10)

	for _, p := range doc.Products {
		// append to pdf
		p.appendColTo(doc)

		if doc.pdf.GetY() > 255 {
			// add page
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
	borderTable := "0"

	// get base Y (top of line)
	baseY := doc.pdf.GetY()

	// name
	doc.pdf.SetX(10)
	doc.pdf.MultiCell(
		90,
		3,
		doc.UnicodeTranslatorFunc(p.Title),
		borderTable,
		"",
		false,
	)

	// compute line height
	colHeight := doc.pdf.GetY() - baseY

	// Product packaging
	if p.Packaging == "" {
		p.Packaging = "---"
	}
	doc.pdf.SetY(baseY)
	doc.pdf.SetX(100)
	doc.pdf.CellFormat(
		20,
		colHeight,
		doc.UnicodeTranslatorFunc(p.Packaging),
		borderTable,
		0,
		"C",
		false,
		0,
		"",
	)

	// unit price
	doc.pdf.SetY(baseY)
	doc.pdf.SetX(120)
	doc.pdf.CellFormat(
		20,
		colHeight,
		doc.UnicodeTranslatorFunc(fmt.Sprintf("%f", p.Price)),
		borderTable,
		0,
		"C",
		false,
		0,
		"",
	)

	// quantity
	doc.pdf.SetX(140)
	doc.pdf.CellFormat(
		20,
		colHeight,
		doc.UnicodeTranslatorFunc(fmt.Sprintf("%f", p.Quantity)),
		borderTable,
		0,
		"C",
		false,
		0,
		"",
	)

	// discount
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
		borderTable,
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
		borderTable,
		0,
		"C",
		false,
		0,
		"",
	)

	// set Y for next line
	doc.pdf.SetY(baseY + colHeight)
}
