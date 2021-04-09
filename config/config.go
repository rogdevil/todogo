package config

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open(
		"postgres", 
		"host=postgrestodo port=5432 user=admin dbname=tododb password=123 sslmode=disabled"
	)
	if err != nil {
		panic(err.Error())
	}

	Db = db

	return DB
}

func GetDB() *gorm.DB {
	return DB
}