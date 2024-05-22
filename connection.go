package connection

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=localhost user=ada-onb0001 password=Kmzway87a@ dbname=belajar_gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
