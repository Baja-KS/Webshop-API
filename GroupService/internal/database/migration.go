package database

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Group{})
	if err != nil {
		return err
	}
	return nil
}
