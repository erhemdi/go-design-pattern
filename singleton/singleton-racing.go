package singleton

import "log"

type DBConnection struct{}

var dbConnection *DBConnection

func GetDBInstance() *DBConnection {
	if dbConnection == nil {
		log.Println("DB Instance created")
		dbConnection = &DBConnection{} // Not thread safe
	}

	return dbConnection
}
