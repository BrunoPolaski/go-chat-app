package mysql

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB
var dbConn *sql.Conn

func Init() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/chat")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err.Error())
	}

	db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, username VARCHAR(255), password VARCHAR(255))")

	database = db
}

func GetConn() *sql.Conn {
	if dbConn == nil {
		conn, err := database.Conn(context.Background())
		if err != nil {
			panic(err.Error())
		}
		dbConn = conn
	}

	return dbConn
}
