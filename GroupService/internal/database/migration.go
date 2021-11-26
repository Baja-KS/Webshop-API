package database

import (
	"gorm.io/gorm"
	"os"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Group{})
	if err != nil {
		return err
	}


	seedTmp:=os.Getenv("SEED_IF_EMPTY")

	seed:=false

	if seedTmp=="true" {
		seed=true
	}

	if !seed {
		return nil
	}
	var count int64
	count=0
	db.Model(&Group{}).Count(&count)

	if count!=0 {
		return nil
	}

	var groups = []Group{
		{
			Name:        "Racunarske komponente",
			Description: "",
		},
		{
			Name:        "Desktop konfiguracije",
			Description: "",
		},
		{
			Name:        "Laptop racunari",
			Description: "",
		},
		{
			Name:        "Mobilni telefoni",
			Description: "",
		},
		{
			Name:        "TV",
			Description: "",
		},
	}
	db.Create(&groups)
	return nil
}
