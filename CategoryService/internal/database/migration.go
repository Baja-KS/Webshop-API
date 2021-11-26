package database

import (
	"gorm.io/gorm"
	"os"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Category{})
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
	db.Model(&Category{}).Count(&count)

	if count!=0 {
		return nil
	}

	var categories = []Category{
		{
			Name:    "Graficke Kartice",
			GroupID: 1,
		},
		{
			Name:    "Procesori",
		},
		{
			Name:    "Maticne ploce",
			GroupID: 1,
		},
		{
			Name:    "Gejmerske konfiguracije",
			GroupID: 2,
		},
		{
			Name:    "Gejmerske polukonfiguracije",
			GroupID: 2,
		},
		{
			Name:    "Gejmerski laptop racunari",
			GroupID: 3,
		},
		{
			Name:    "Ostali laptop racunari",
			GroupID: 3,
		},
	}
	db.Create(&categories)
	return nil
}
