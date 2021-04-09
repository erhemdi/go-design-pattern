package singleton

import (
	"log"
	"sync"
)

var once sync.Once

func GetDBInstanceSync() *DBConnection {
	once.Do(func() {
		if dbConnection == nil {
			log.Println("DB Instance created")
			dbConnection = &DBConnection{}
		}
	})

	return dbConnection
}
