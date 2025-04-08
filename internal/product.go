package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type ProductsController struct {
	Prod []Product
	n    int
}

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is-published,omitempty"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func LoadSliceProducts() (ProductsController, error) {
	p := ProductsController{}
	file, err := os.Open("products.json")
	if err != nil {
		return p, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	reader, err := io.ReadAll(file)
	if err != nil {
		return p, fmt.Errorf("error reading file: %v", err)
	}
	json.Unmarshal(([]byte(reader)), &p.Prod)
	p.n = len(p.Prod)
	return p, nil
}

func (p *Product) validateProduct() error {
	if p.Name == "" {
		return fmt.Errorf("name is required")
	}
	if p.Quantity <= 0 {
		return fmt.Errorf("quantity must be greater than 0")
	}
	if p.CodeValue == "" {
		return fmt.Errorf("code value is required")
	}
	if p.Expiration == "" {
		return fmt.Errorf("expiration is required")
	}
	if p.Price <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}
	if !validDate(p.Expiration) {
		return fmt.Errorf("expiration date is not valid")
	}
	return nil
}

func validDate(date string) bool {
	_, err := time.Parse("01-02-2025", date)
	if err != nil {
		return false
	}
	return true
}

func (p *ProductsController) AddProduct(prod Product) error {
	err := prod.validateProduct()
	if err != nil {
		return err
	}
	err = p.codeInUse(prod.CodeValue)
	if err != nil {
		return err
	}
	prod.ID = p.n + 1
	p.Prod = append(p.Prod, prod)
	p.n++
	return nil
}

func (p *ProductsController) codeInUse(code string) error {
	for _, product := range p.Prod {
		if product.CodeValue == code {
			return fmt.Errorf("code value is already in use")
		}
	}
	return nil
}
