package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name         string `gorm:"type:varchar(100)" json:"name"`
	Email        string `gorm:"type:varchar(100)" json:"email"`
	Address      string `gorm:"type:varchar(100)" json:"address"`
	PasswordHash string `gorm:"type:varchar(100)" json:"password_hash"`
}
