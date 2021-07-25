package database

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Product{})
	if err != nil {
		return err
	}
	return nil
}
