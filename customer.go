package goinvoice

import "fmt"

func (doc *Document) writeCustomer() {
	doc.pdf.SetFont(doc.Font, "", 12)

	const align = "R"

	doc.pdf.SetXY(130, 20)

	// Name rect
	doc.pdf.Rect(130, doc.pdf.GetY(), 70, 10, "F")

	// Set name
	doc.pdf.CellFormat(70, 10, doc.UnicodeTranslatorFunc(doc.customerNameToString()), "0", 0, align, false, 0, "")

	// Info rect
	var infoRectHeight float64 = 4
	var startY float64
	if doc.Customer.DeliveryType != "" {
		infoRectHeight += 5
		startY += 3.5
	}
	if doc.Customer.PhoneNumber != "" {
		infoRectHeight += 5
		startY += 3.5
	}
	if doc.Customer.Email != "" {
		infoRectHeight += 5
		startY += 3.5
	}
	if doc.Customer.DeliveryInfo != "" {
		infoRectHeight += 5
		startY += 3.5
	}

	if infoRectHeight > 4 {
		doc.pdf.Rect(130, doc.pdf.GetY()+12, 70, infoRectHeight, "F")
	}

	// Set address
	doc.pdf.SetXY(130, doc.pdf.GetY()+14)
	doc.pdf.MultiCell(70, 5, doc.UnicodeTranslatorFunc(doc.customerInfoToString()), "0", align, false)
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
