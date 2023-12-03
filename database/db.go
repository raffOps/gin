package database

import (
	"github.com/rjribeiro/goBackend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDB() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable TimeZone=UTC"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Nao foi possivel conectar o banco de dados")
	}
	err := DB.AutoMigrate(&models.Aluno{})
	if err != nil {
		return
	}
}
