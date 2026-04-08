package migrations

import (
	"fmt"
	models "transaction/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	fmt.Println("Running migrations...")

	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	err := db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.ProductHistory{},
		&models.Customer{},
	)

	if err != nil {
		return err
	}

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	fmt.Println("Migration success")
	return nil
}

func DropAll(db *gorm.DB) error {
	fmt.Println("Dropping tables...")

	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	err := db.Migrator().DropTable(
		&models.ProductHistory{},
		&models.Product{},
		&models.Category{},
		&models.User{},
		&models.Customer{},
	)

	if err != nil {
		return err
	}

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	fmt.Println("Drop success")
	return nil
}
