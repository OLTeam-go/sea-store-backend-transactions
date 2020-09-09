package database

import (
	"fmt"
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var singleton *gorm.DB
var mt sync.Mutex

func connectDatabase() (*gorm.DB, error) {
	_ = godotenv.Load()

	dbURL, exist := os.LookupEnv("DATABASE_URL")
	if !exist {
		panic("DATABASE_URL did not exists")
	}
	fmt.Println(fmt.Sprintf("connceting to postgres = %s", dbURL))

	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "db connection failed: %s", err.Error())
	}
	if err != nil {
		return nil, err
	}
	return db, nil
}

// GetInstance return the singleton of db connection
func GetInstance() (*gorm.DB, error) {
	if singleton == nil {
		mt.Lock()
		defer mt.Unlock()
		if singleton == nil {
			db, err := connectDatabase()
			if err != nil {
				return nil, err
			}
			singleton = db
			return singleton, nil
		}
	}
	return singleton, nil
}
