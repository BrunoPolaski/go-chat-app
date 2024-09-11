package contract

import (
	"database/sql"

	"github.com/BrunoPolaski/go-chat-app/pkg/utility"
)

type Database interface {
	Init() *utility.RestErr
	GetConn() (*sql.Conn, *utility.RestErr)
}
