package database

import (
	"errors"
	"fmt"

	"github.com/pyramid.io/planit-backend/pkg/framework/database/drivers"
	"github.com/pyramid.io/planit-backend/pkg/framework/interfaces"
)

type DatabaseServiceInterface interface {
	AddConnection(name string, connection drivers.DatabaseDriverInterface) error
	GetConnection(name string) (drivers.DatabaseDriverInterface, error)
}

type DatabaseService struct {
	Connections map[string]drivers.DatabaseDriverInterface
}

func (db *DatabaseService) AddConnection(name string, connection drivers.DatabaseDriverInterface) error {
	db.Connections[name] = connection
	return nil
}

func (db *DatabaseService) GetConnection(name string) (drivers.DatabaseDriverInterface, error) {
	connection, exists := db.Connections[name]
	if !exists {
		return nil, fmt.Errorf("connection %s not found", name)
	}
	return connection, nil
}

var dbService DatabaseServiceInterface = &DatabaseService{
	Connections: make(map[string]drivers.DatabaseDriverInterface),
}

func New(connectionsConfig []interfaces.DatabaseConnectionConfigInterface) (DatabaseServiceInterface, error) {
	for _, connectionConfig := range connectionsConfig {
		var constructor drivers.DatabaseConnectionConstructor
		var ok bool

		switch v := connectionConfig.GetDriver().(type) {
		case string:
			constructor, ok = drivers.ServiceRegistry[v]
			if !ok {
				return nil, errors.New("database connection is not found in registry")
			}
		case drivers.DatabaseConnectionConstructor:
			constructor = v
		default:
			return nil, errors.New("invalid connection key or constructor")
		}

		connection, err := constructor(connectionConfig.GetConfig())
		if err != nil {
			return nil, err
		}

		dbService.AddConnection(
			connectionConfig.GetConnectionName(),
			connection,
		)
	}
	return dbService, nil
}
