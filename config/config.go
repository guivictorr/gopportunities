package config

import "gorm.io/gorm"

var db *gorm.DB

func Init() error {
	return nil
}
