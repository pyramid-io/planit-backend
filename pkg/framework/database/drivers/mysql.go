package drivers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pyramid.io/planit-backend/pkg/framework/database/db_result"
)

type MysqlDriver struct {
	dsn string
	config map[string]interface{}
	connection *sql.DB
}

func (driver *MysqlDriver) Select(statement string, queryParams ...any) (*db_result.SelectResult, error) {
	dbRows, err := driver.connection.Query(statement, queryParams...)
	if (err != nil) {
		return nil, err
	}
	
	defer dbRows.Close()

	columns, err := dbRows.Columns()
    if err != nil {
        log.Fatal(err)
    }

    values := make([]interface{}, len(columns))
    for i := range values {
        var value interface{}
        values[i] = &value
    }

    var result db_result.SelectResult

    for dbRows.Next() {
        err := dbRows.Scan(values...)
        if err != nil {
            log.Fatal(err)
        }

        record := make(db_result.Row)
        for i, colName := range columns {
            rawValue := *(values[i].(*interface{}))

            if rawValue != nil {
                switch v := rawValue.(type) {
                case []byte:
					record[colName] = string(v) 
                default:
                    record[colName] = v
                }
            } else {
                record[colName] = nil
            }
        }
		result.Add(record)
    }

	return &result, nil
}

func (driver *MysqlDriver) Exec(statement string, queryParams ...any) (*db_result.ExecuteResult, error) {
	Result, err := driver.connection.Exec(statement, queryParams...)

	if err != nil {
		log.Fatalf("error while insert to database: ", err)
	}

	id, err := Result.LastInsertId()
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id: %s", err)
	}

	count, err := Result.RowsAffected()
	if err != nil {
		log.Fatalf("impossible to retrieve affected rows: ", err)
	}
	
	return &db_result.ExecuteResult{
		LastInsertId: id,
		RowsAffected: count,
	}, nil
}

func (driver *MysqlDriver) Connect() error {

    connection, err := sql.Open("mysql", driver.dsn)
    if err != nil {
        panic(err)
    }

	if maxOpenConnections, exists := driver.config["maxOpenConnections"]; exists {
		val, ok := maxOpenConnections.(int)
		if !ok{
			fmt.Println("max open connections in config must be defined as int")
		} else {
			connection.SetMaxOpenConns(val)
		}
	}

	if maxLifeTime, exists := driver.config["maxLifeTime"]; exists {
		val, ok := maxLifeTime.(time.Duration)
		if !ok{
			fmt.Println("max life time in config must be time.Duration")
		} else {
			connection.SetConnMaxLifetime(val)
		}
	}

	if macIdleConnection, exists := driver.config["maxIdleConnections"]; exists {
		val, ok := macIdleConnection.(int)
		if !ok{
			fmt.Println("max open connections in config must be int")
		} else {
			connection.SetMaxIdleConns(val)
		}
	}

	driver.connection = connection

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
	charset, charsetExists := config["charset"].(string)

	if (!usernameExists || !passwordExists || !hostExists || !portExists || !databaseNameExists) {
		return nil, errors.New("mysql connector config is missing")
	}

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + databaseName

	if charsetExists {	
		dsn = dsn + "?" + "charset=" + charset
	}

	return &dsn, nil
}
