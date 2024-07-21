package drivers

import "github.com/pyramid.io/planit-backend/pkg/framework/database/db_result"

type DatabaseDriverInterface interface {
	Connect() error
	Close() error
	Select(statement string, queryParams ...any) (*db_result.SelectResult, error)
	Exec(statement string, queryParams ...any) (*db_result.ExecuteResult, error)
}