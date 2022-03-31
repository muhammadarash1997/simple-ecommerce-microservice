package user

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	GetByUUID(id string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (this *repository) Save(user User) (User, error) {
	_, err := this.FindByEmail(user.Email)
	if err == nil {
		return user, errors.New("email has been used")
	}

	err = this.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (this *repository) FindByEmail(email string) (User, error) {
	user := User{}

	err := this.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, errors.New("Email has not been registered")
	}

	return user, nil
}

func (this *repository) GetByUUID(id string) (User, error) {
	user := User{}

	err := this.db.First(&user, "id = ?", id).Error
	if err != nil {
		return user, errors.New("User not found")
	}

	return user, nil
}
