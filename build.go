package goinvoice

import (
	"github.com/jung-kurt/gofpdf"
)

func (doc *Document) BuildPdf() (gofpdf.Fpdf, error) {
	if doc.Company != nil {
		doc.writeCompany()
	}

	if doc.Customer != nil {
		doc.writeCustomer()
	}

	doc.drawsTableTitles()

	if doc.Products != nil {
		doc.writeProducts()
	}

	doc.appendTotal()
	return *doc.pdf, nil
}
