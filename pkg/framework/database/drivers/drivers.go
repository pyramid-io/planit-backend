package drivers

type DatabaseDriverInterface interface {
	Connect() error
	Close() error
}