package models

import "errors"

type Category struct {
	Model
	Description string     `gorm:"size:512;not null;unique" json:"description"`
	Product     []*Product `gorm:"foreignkey:CategoryID" json:"products"`
}

var (
	ErrCategoryEmptyDescription = errors.New("Description name cant be empty")
)

func (c *Category) Validate() error {
	if c.Description == "" {
		return ErrCategoryEmptyDescription
	}
	return nil
}
