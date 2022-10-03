package main

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

	return *doc.pdf, nil
}
