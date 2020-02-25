package models

type User struct {
	ID        uint32 `gorm:"primary_key; auto_increment" json:"id"`
	Nickname  string `gorm:"size:20; not null; unique" json:"nickname"`
	Email     string `gorm:"size:50; not null; unique" json:"email"`
	Password  string `gorm:"size:20; not null; unique" json:"user_name"`
	CreatedAt string `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt string `gorm:"default:current_timestamp()" json:"updated_at"`
}
