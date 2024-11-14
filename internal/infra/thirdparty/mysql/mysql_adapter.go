package mysql

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/BrunoPolaski/go-chat-app/pkg/utility"
	_ "github.com/go-sql-driver/mysql"
)

type mySQLAdapter struct {
	db *sql.DB
}

func NewMySQLAdapter() *mySQLAdapter {
	return &mySQLAdapter{}
}

func (msa *mySQLAdapter) Connect() (*sql.DB, error) {
	if msa.db != nil {
		return msa.db, nil
	}

	var err error
	msa.db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/chat")
	if err != nil {
		return nil, err
	}

	if err = msa.db.Ping(); err != nil {
		return nil, err
	}

	return msa.db, nil
}

func (msa *mySQLAdapter) Disconnect() *utility.RestErr {
	if msa.db != nil {
		if err := msa.db.Close(); err != nil {
			return utility.NewInternalServerError("error when trying to disconnect from database")
		}
	}
	return nil
}

func (msa *mySQLAdapter) SelectAll(table string, params []string, ascendingOrderBy bool) (*sql.Rows, error) {
	query := "SELECT * FROM " + table

	if len(params) > 0 {
		query += " WHERE "
		for i, param := range params {
			if i > 0 {
				query += " AND "
			}
			query += param
		}
	}

	if ascendingOrderBy {
		query += " ORDER BY id ASC"
	} else {
		query += " ORDER BY id DESC"
	}

	return msa.db.Query(query)
}

func (msa *mySQLAdapter) Insert(table string, fields, values []string) (sql.Result, error) {
	if len(fields) != len(values) {
		return nil, errors.New("fields and values must have the same length")
	}
	msa.db.Begin()
	query, err := msa.db.Prepare("INSERT INTO " + table + " (" + strings.Join(fields, ",") + ") VALUES (" + strings.Repeat("?,", len(values)-1) + "?)")
	if err != nil {
		return nil, err
	}
	defer query.Close()

	return query.Exec(values)
}

func (msa *mySQLAdapter) Update(table string, fields, values, conditions []string) (sql.Result, error) {
	if len(fields) != len(values) {
		return nil, errors.New("fields and values must have the same length")
	}
	msa.db.Begin()
	query, err := msa.db.Prepare("UPDATE " + table + " SET " + strings.Join(fields, "=?, ") + "=? WHERE " + strings.Join(conditions, " AND "))
	if err != nil {
		return nil, err
	}
	defer query.Close()

	return query.Exec(values)
}

func (msa *mySQLAdapter) Delete(table string, conditions []string) (sql.Result, error) {
	msa.db.Begin()
	query, err := msa.db.Prepare("DELETE FROM " + table + " WHERE " + strings.Join(conditions, " AND "))
	if err != nil {
		return nil, err
	}
	defer query.Close()

	return query.Exec()
}
