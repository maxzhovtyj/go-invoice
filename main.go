package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	doc, err := NewWithCyrillic(pwd)
	if err != nil {
		return
	}

	doc.setPwd(pwd)
	doc.SetInvoice(&Invoice{
		Logo:  nil,
		Title: "Alliance Cup",
	})
	doc.SetCompany(&Company{
		Name:        "AllianceCup",
		Address:     "Шухевича, 22",
		PhoneNumber: "+380(96) 512-15-16",
		City:        "Рівне",
		Country:     "Україна",
	})
	doc.SetCustomer(&Customer{
		LastName:     "Жовтанюк",
		FirstName:    "Максим",
		MiddleName:   "В'ячеславович",
		PhoneNumber:  "+380(68) 306-29-75",
		Email:        "zhovtyjshady@gmail.com",
		DeliveryType: "Доставка новою поштою",
		DeliveryInfo: "м. Рівне, Відділення №12",
	})

	pdf, err := doc.BuildPdf()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = pdf.OutputFileAndClose("output.pdf")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

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

func (doc *Document) setPwd(pwd string) {
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
