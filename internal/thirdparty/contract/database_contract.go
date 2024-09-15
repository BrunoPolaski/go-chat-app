package contract

type DatabaseContract interface {
	Connect() error
	Disconnect() error
}
