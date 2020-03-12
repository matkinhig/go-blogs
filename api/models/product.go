package models

import "errors"

type ProductStatus uint16

const (
	ProductStatus_Unavailable = 0
	ProductStatus_Available   = 1
)

type Product struct {
	Model
	Name    string        `gorm:"size:512; not null" json:"name"`
	Price   float64       `gorm:"type:decimal(10,2);not null;default:0.0" json:"price"`
	Quality uint16        `gorm:"default:0;unsigned" json:"quantity"`
	Status  ProductStatus `gorm:"char(1);default:1" json:"status"`
}

var (
	ErrProductEmptyName = errors.New("Product name cant be empty")
)

func (p *Product) Validate() error {
	if p.Name == "" {
		return ErrProductEmptyName
	}
	return nil
}
