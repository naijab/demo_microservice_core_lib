package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

var dbType = os.Getenv("DB_TYPE")
var dbConnectionString = os.Getenv("DB_CONNECTION_STRING")

type DatabaseConnection struct {
	Database *gorm.DB
}

func (h *DatabaseConnection) NewDatabaseConnection() error {
	db, err := gorm.Open(dbType, dbConnectionString)
	if err != nil {
		log.Print("Connection DB Error => ", err)
		return err
	}
	h.Database = db
	log.Print("Connection DB Completed")
	return nil
}
