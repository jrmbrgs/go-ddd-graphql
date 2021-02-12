package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/jrmbrgs/goodkarma-api/library/models"
)

type dbConfig struct {
	host     string
	port     int
	user     string
	dbname   string
	password string
}

var pwd string = os.Getenv("GK_DB_PASS")

var config = dbConfig{"localhost", 5432, "gk", "goodkarma", pwd}

func getDatabaseUrl() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s",
		config.host, config.port, config.user, config.dbname, config.password)
}

func GetDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  getDatabaseUrl(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	return db, err
}

func RunMigrations(db *gorm.DB) {
	log.Println("Running migration")

	db.Migrator().AutoMigrate(&models.OffenseCategory{})
	db.Migrator().AutoMigrate(&models.Offense{})
	db.Migrator().AutoMigrate(&models.Human{})
	db.Migrator().AutoMigrate(&models.CharityOrganization{})
	db.Migrator().AutoMigrate(&models.CharityOrganizationTag{})
}
