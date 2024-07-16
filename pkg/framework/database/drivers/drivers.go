package drivers

import "github.com/pyramid.io/planit-backend/pkg/framework/database/result"

type DatabaseDriverInterface interface {
	Connect() error
	Close() error
	Select(statement string, queryParams ...any) (*result.RowCollection, error)
}