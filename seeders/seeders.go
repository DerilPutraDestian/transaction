package seeders

import (
	"fmt"
	models "transaction/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	fmt.Println("Seeding database...")

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)

	user := models.User{
		Name:     "Deril Admin",
		Email:    "deril@mail.com",
		Password: string(hashedPassword),
		Role:     "admin",
	}

	if err := db.Where(models.User{Email: "deril@mail.com"}).FirstOrCreate(&user).Error; err != nil {
		return err
	}

	fmt.Println("Seeding success!")
	return nil
}
