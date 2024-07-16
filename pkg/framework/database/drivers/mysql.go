package drivers

type MysqlDriver struct {}

func (mysql *MysqlDriver) Connect() error {
	return nil
}
func (mysql *MysqlDriver) Close() error {
	return nil
}

func MysqlConnectionConstructor(config map[string]interface{}) (DatabaseDriverInterface, error) {
	return &MysqlDriver{}, nil
}
