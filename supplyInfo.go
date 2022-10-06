package goinvoice

func (doc *Document) supplierAndReceiver() float64 {
	var supplierTitle, recipientTitle string

	switch doc.Language {
	case "UK":
		supplierTitle = "Постачальник:"
		recipientTitle = "Отримав:"
	default:
		supplierTitle = "Supplier:"
		recipientTitle = "Recipient:"
	}

	doc.pdf.SetFont(doc.Font, "", 10)

	posY := doc.pdf.GetY()

	// draw supplier text
	doc.pdf.SetXY(10, posY+5)
	doc.pdf.CellFormat(26, 5, doc.UnicodeTranslatorFunc(supplierTitle), "0", 0, "R", false, 0, "")

	// draw supplier line
	doc.pdf.SetXY(38, posY+8.5)
	doc.pdf.MultiCell(70, 0, doc.UnicodeTranslatorFunc(""), "B", "L", false)

	// draw receiver text
	doc.pdf.SetXY(10, posY+10)
	doc.pdf.CellFormat(26, 5, doc.UnicodeTranslatorFunc(recipientTitle), "0", 0, "R", false, 0, "")

	// draw receiver line
	doc.pdf.SetXY(38, posY+13.5)
	doc.pdf.MultiCell(70, 0, doc.UnicodeTranslatorFunc(""), "B", "L", false)
	return posY
}
