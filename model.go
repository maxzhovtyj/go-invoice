package main

import "github.com/jung-kurt/gofpdf"

type Invoice struct {
	Logo  []byte
	Title string
}

type Company struct {
	Name        string
	Address     string
	PhoneNumber string
	City        string
	Country     string
}

type Customer struct {
	LastName     string
	FirstName    string
	MiddleName   string
	PhoneNumber  string
	Email        string
	DeliveryType string
	DeliveryInfo string
}

type Product struct {
	Title    string
	Price    string
	Quantity string
	Discount string
	Total    string
}

type Document struct {
	UnicodeTranslatorFunc func(string) string
	Products              []Product
	*Customer
	*Company
	*Invoice
	Pwd string
	pdf *gofpdf.Fpdf
}
