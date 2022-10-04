package goinvoice

import "fmt"

func (doc *Document) writeCompany() {
	// set X and Y
	doc.pdf.SetXY(10, 20)

	// draw rect
	doc.pdf.SetFillColor(212, 212, 212)
	doc.pdf.Rect(10, 20, 70, 10, "F")

	// Draw text
	doc.pdf.SetFont("Helvetica", "B", 12)
	doc.pdf.CellFormat(80, 10, doc.UnicodeTranslatorFunc(doc.Company.Name), "0", 0, "L", false, 0, "")

	// draw address rect
	doc.pdf.Rect(10, doc.pdf.GetY()+12, 70, 19, "F")

	// Set address
	doc.pdf.SetFont("Helvetica", "", 12)
	doc.pdf.SetXY(10, doc.pdf.GetY()+14)
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
