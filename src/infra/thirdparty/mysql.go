package thirdparty

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

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
}

func GetConnection() *sql.Conn {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/chat")
	if err != nil {
		panic(err.Error())
	}

	if conn, err := db.Conn(context.Background()); err != nil {
		panic(err.Error())
	} else {
		return conn
	}
}
