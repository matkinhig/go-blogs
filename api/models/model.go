package models

import "time"

//code next in here
type Model struct {
	ID         uint32 `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"default:current_timestamp;not null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:current_timestamp;not null" json:"updated_at"`
}
