package mssql

import (
	"database/sql"
)

//DB - database
var DB *sql.DB

//NewConnection - new connection
func NewConnection() {
}

//GetConnection - get connection
func GetConnection() *sql.DB {
	return DB
}
