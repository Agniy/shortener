package storage

import (
	"fmt"
	"github.com/Agniy/shortener/internal/app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var dbInstance *gorm.DB

// Used during creation of singleton client object in GetDbClient().
var dbInstanceError error

// Used to execute client creation procedure only once.
var dbOnce sync.Once

func NewPostgresConnection() (*gorm.DB, error) {
	cfg := config.GetConfig()
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.Dbname)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	// ping db
	psqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	err = psqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
func GetDbClient() (*gorm.DB, error) {
	dbOnce.Do(func() {
		db, err := NewPostgresConnection()
		if err != nil {
			dbInstanceError = err
		}
		dbInstance = db
	})
	return dbInstance, dbInstanceError
}
