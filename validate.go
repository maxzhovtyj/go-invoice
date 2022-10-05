package goinvoice

import "fmt"

func (doc *Document) validate() error {
	// validate gofpdf instance
	if doc.pdf == nil {
		return fmt.Errorf("failed to create pdf, pdf instance wasn't initialized")
	}

	// validate company info
	if doc.Company != nil {
		return fmt.Errorf("failed to create pdf, company wasn't provided")
	}
	if doc.Company.Name == "" {
		return fmt.Errorf("failed to create pdf, company name wasn't provided")
	}
	if doc.Company.Address == "" {
		return fmt.Errorf("failed to create pdf, company address wasn't provided")
	}
	if doc.Company.PhoneNumber == "" {
		return fmt.Errorf("failed to create pdf, company phone number wasn't provided")
	}
	if doc.Company.City == "" {
		return fmt.Errorf("failed to create pdf, company city wasn't provided")
	}
	if doc.Company.Country == "" {
		return fmt.Errorf("failed to create pdf, company country wasn't provided")
	}

	// validate customer
	if doc.Customer != nil {
		return fmt.Errorf("failed to create pdf, customer wasn't provided")
	}
	if doc.Customer.LastName != "" {
		return fmt.Errorf("failed to create pdf, customer lastname wasn't provided")
	}
	if doc.Customer.FirstName != "" {
		return fmt.Errorf("failed to create pdf, customer firstname wasn't provided")
	}
	if doc.Customer.MiddleName != "" {
		return fmt.Errorf("failed to create pdf, customer middlename wasn't provided")
	}

	// validate products
	if doc.Products == nil || len(doc.Products) == 0 {
		return fmt.Errorf("failed to create pdf, products wasn't provided")
	}

	return nil
}
