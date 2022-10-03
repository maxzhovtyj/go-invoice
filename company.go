package main

import "fmt"

func (doc *Document) writeCompany() {
	// set X and Y
	doc.pdf.SetXY(20, 20)

	// draw rect
	doc.pdf.SetFillColor(212, 212, 212)
	doc.pdf.Rect(20, 20, 70, 10, "F")

	// Draw text
	doc.pdf.SetFont("Helvetica", "B", 12)
	doc.pdf.CellFormat(80, 10, doc.UnicodeTranslatorFunc(doc.Company.Name), "0", 0, "L", false, 0, "")

	doc.pdf.Rect(20, doc.pdf.GetY()+12, 70, 17, "F")

	// Set address
	doc.pdf.SetFont("Helvetica", "", 12)
	doc.pdf.SetXY(20, doc.pdf.GetY()+13)
	doc.pdf.MultiCell(80, 5, doc.UnicodeTranslatorFunc(doc.companyToString()), "0", "L", false)
}

func (doc *Document) companyToString() string {
	return fmt.Sprintf(
		"%s\n%s\n%s, %s",
		doc.Company.Address,
		doc.Company.PhoneNumber,
		doc.Company.City,
		doc.Company.Country,
	)
}
