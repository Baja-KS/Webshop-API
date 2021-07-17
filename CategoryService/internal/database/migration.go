package database

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Category{})
	if err != nil {
		return err
	}
	return nil
}
