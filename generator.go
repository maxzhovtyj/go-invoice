package goinvoice

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

func NewWithCyrillic(pwd string) (*Document, error) {
	tr, err := gofpdf.UnicodeTranslatorFromFile(pwd + "/font/" + "cp1251.map")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var doc Document
	doc.pdf = gofpdf.New("P", "mm", "A4", pwd+"/font/")

	doc.SetUnicodeTranslator(tr)
	doc.pdf.AddFont("Helvetica", "", "helvetica_1251.json")
	doc.pdf.AddPage()

	return &doc, err
}

func New(pwd string) (*Document, error) {
	var doc Document
	doc.pdf = gofpdf.New("P", "mm", "A4", pwd+"/font/")
	return &doc, nil
}

func (doc *Document) SetPwd(pwd string) {
	doc.Pwd = pwd
}

func (doc *Document) SetUnicodeTranslator(f func(string) string) {
	doc.UnicodeTranslatorFunc = f
}

func (doc *Document) SetInvoice(invoice *Invoice) {
	doc.Invoice = invoice
}

func (doc *Document) SetCompany(company *Company) {
	doc.Company = company
}

func (doc *Document) SetCustomer(customer *Customer) {
	doc.Customer = customer
}

func (doc *Document) AppendProductItem(product *Product) {
	doc.Products = append(doc.Products, *product)
}
