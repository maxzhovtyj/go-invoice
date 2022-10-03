package main

import "fmt"

func (doc *Document) writeCustomer() {
	doc.pdf.SetXY(130, 20)

	// Name rect
	doc.pdf.Rect(130, doc.pdf.GetY(), 70, 10, "F")

	// Set name
	doc.pdf.SetFont("Helvetica", "", 12)
	doc.pdf.Cell(40, 10, doc.UnicodeTranslatorFunc(doc.customerNameToString()))
	doc.pdf.SetFont("Helvetica", "", 12)

	// Address rect
	var addrRectHeight float64 = 22

	doc.pdf.Rect(130, doc.pdf.GetY()+12, 70, addrRectHeight, "F")

	// Set address
	doc.pdf.SetFont("Helvetica", "", 12)
	doc.pdf.SetXY(130, doc.pdf.GetY()+13)
	doc.pdf.MultiCell(70, 5, doc.UnicodeTranslatorFunc(doc.customerInfoToString()), "0", "L", false)
}

func (doc *Document) customerInfoToString() string {
	return fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		doc.Customer.PhoneNumber,
		doc.Customer.Email,
		doc.Customer.DeliveryType,
		doc.Customer.DeliveryInfo,
	)
}

func (doc *Document) customerNameToString() string {
	return fmt.Sprintf(
		"%s %s.%s",
		doc.Customer.LastName,
		string([]rune(doc.Customer.FirstName)[0]),
		string([]rune(doc.Customer.MiddleName)[0]),
	)
}
