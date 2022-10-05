package goinvoice

// drawsTableTitles in document
func (doc *Document) drawsTableTitles() {
	// Draw table titles
	doc.pdf.SetX(10)
	doc.pdf.SetY(doc.pdf.GetY() + 5)
	doc.pdf.SetFont("Helvetica", "", 10)

	// Draw rec
	doc.pdf.SetFillColor(212, 212, 212)
	doc.pdf.Rect(10, doc.pdf.GetY(), 190, 6, "F")

	// Name
	doc.pdf.SetX(10)
	doc.pdf.CellFormat(
		90,
		6,
		doc.UnicodeTranslatorFunc("Товар"),
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
		doc.UnicodeTranslatorFunc("Пакування"),
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
		doc.UnicodeTranslatorFunc("Ціна"),
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
		doc.UnicodeTranslatorFunc("Кількість"),
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
		doc.UnicodeTranslatorFunc("Знижка"),
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
		doc.UnicodeTranslatorFunc("Сума"),
		"0",
		0,
		"C",
		false,
		0,
		"",
	)
}
