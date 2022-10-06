package goinvoice

// drawsTableTitles in document
func (doc *Document) drawsTableTitles() {
	var (
		productTitle,
		packagingTitle,
		priceTitle,
		quantityTitle,
		discountTitle,
		totalTitle string
	)

	switch doc.Language {
	case "UK":
		productTitle = "Товар"
		packagingTitle = "Пакування"
		priceTitle = "Ціна"
		quantityTitle = "Кількість"
		discountTitle = "Знижка"
		totalTitle = "Сума"
	default:
		productTitle = "Name"
		packagingTitle = "Packaging"
		priceTitle = "Price"
		quantityTitle = "Quantity"
		discountTitle = "Discount"
		totalTitle = "Total"
	}

	doc.pdf.SetFont(doc.Font, "", 10)

	// Draw table titles
	doc.pdf.SetX(10)
	doc.pdf.SetY(doc.pdf.GetY() + 5)

	// Draw rec
	doc.pdf.SetFillColor(212, 212, 212)
	doc.pdf.Rect(10, doc.pdf.GetY(), 190, 6, "F")

	// Name
	doc.pdf.SetX(10)
	doc.pdf.CellFormat(
		90,
		6,
		doc.UnicodeTranslatorFunc(productTitle),
		"0",
		0,
		"",
		false,
		0,
		"",
	)

	doc.pdf.SetX(100)
	doc.pdf.CellFormat(
		20,
		6,
		doc.UnicodeTranslatorFunc(packagingTitle),
		"0",
		0,
		"C",
		false,
		0,
		"",
	)

	// Unit price
	doc.pdf.SetX(120)
	doc.pdf.CellFormat(
		20,
		6,
		doc.UnicodeTranslatorFunc(priceTitle),
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
		6,
		doc.UnicodeTranslatorFunc(quantityTitle),
		"0",
		0,
		"C",
		false,
		0,
		"",
	)

	// Discount
	doc.pdf.SetX(160)
	doc.pdf.CellFormat(
		20,
		6,
		doc.UnicodeTranslatorFunc(discountTitle),
		"0",
		0,
		"C",
		false,
		0,
		"",
	)

	// Total
	doc.pdf.SetX(180)
	doc.pdf.CellFormat(
		20,
		6,
		doc.UnicodeTranslatorFunc(totalTitle),
		"0",
		0,
		"C",
		false,
		0,
		"",
	)
}
