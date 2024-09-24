package contract

import "database/sql"

type DatabaseContract interface {
	Connect() (*sql.DB, error)
	Disconnect() error
}
