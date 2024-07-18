package drivers

type DatabaseConnectionConstructor func(config map[string]interface{}) (DatabaseDriverInterface, error)

var ServiceRegistry = map[string]DatabaseConnectionConstructor{
	"mysql": MysqlConnectionConstructor,
}
