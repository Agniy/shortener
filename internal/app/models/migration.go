package models

import "gorm.io/gorm"

func MigrateAllModels(db *gorm.DB) error {
	err := db.AutoMigrate(&Links{})
	if err != nil {
		return err
	}
	return nil
}
