package drivers

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/pyramid.io/planit-backend/pkg/framework/database/result"
)

type MysqlDriver struct {
	dsn string
	config map[string]interface{}
	connection *sql.DB
}

func (mysqlDriver *MysqlDriver) Select(statement string, queryParams ...any) (*result.RowCollection, error) {
	rows, err := mysqlDriver.connection.Query(statement, queryParams...)

	if (err != nil) {
		return nil, err
	}

	rowCollection := &result.RowCollection{}
	columns, err := rows.Columns()
    if err != nil {
        log.Fatal(err)
    }

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
        for i := range columns {
            valuePtrs[i] = &values[i]
        }

		err := rows.Scan(valuePtrs...)
        if err != nil {
            log.Fatal(err)
        }

		row := make(result.Row)
        for i, colName := range columns {
            row[colName] = values[i]
        }

		rowCollection.Add(row)
	}

	return rowCollection, nil
}

func (mysqlDriver *MysqlDriver) Connect() error {

    connection, err := sql.Open("mysql", mysqlDriver.dsn)
    if err != nil {
        panic(err)
    }

	if maxOpenConnections, exists := mysqlDriver.config["maxOpenConnections"]; exists {
		val, ok := maxOpenConnections.(int)
		if !ok{
			fmt.Println("max open connections in config must be defined as int")
		} else {
			connection.SetMaxOpenConns(val)
		}
	}

	if maxLifeTime, exists := mysqlDriver.config["maxLifeTime"]; exists {
		val, ok := maxLifeTime.(time.Duration)
		if !ok{
			fmt.Println("max life time in config must be time.Duration")
		} else {
			connection.SetConnMaxLifetime(val)
		}
	}

	if macIdleConnection, exists := mysqlDriver.config["maxIdleConnections"]; exists {
		val, ok := macIdleConnection.(int)
		if !ok{
			fmt.Println("max open connections in config must be int")
		} else {
			connection.SetMaxIdleConns(val)
		}
	}

	mysqlDriver.connection = connection

	return nil
}

func (mysqlDriver *MysqlDriver) Close() error {
	
	err := mysqlDriver.connection.Close()
	
	if (err != nil){
		return err
	}

	return nil

}

func MysqlConnectionConstructor(config map[string]interface{}) (DatabaseDriverInterface, error) {
	dsn, error := getDsnByConfig(config)
	if (error != nil) {
		return nil, error
	}

	return &MysqlDriver{
		dsn : *dsn,
		config :config,
	}, nil
}

func getDsnByConfig(config map[string]interface{}) (*string, error) {
	username, usernameExists := config["username"].(string)
	password, passwordExists := config["password"].(string)
	host, hostExists := config["host"].(string)
	port, portExists := config["port"].(string)
	databaseName, databaseNameExists := config["databaseName"].(string)

	if (!usernameExists || !passwordExists || !hostExists || !portExists || !databaseNameExists) {
		return nil, errors.New("mysql connector config is missing")
	}

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + databaseName
	
	return &dsn, nil
}
