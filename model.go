package goinvoice

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
	Title     string
	Packaging string
	Price     float64
	Quantity  float64
	Discount  float64
	Total     float64
}

type Document struct {
	UnicodeTranslatorFunc func(string) string
	Products              []Product
	OrderId               string
	*Invoice
	*Customer
	*Company
	Language string
	Font     string
	TotalSum float64
	Pwd      string
	pdf      *gofpdf.Fpdf
}
