package db_result

type ExecuteResult struct {
	LastInsertId int64
	RowsAffected int64
}

func (result *ExecuteResult) GetLastInsertId() (int64, error){
	return result.LastInsertId, nil
}

func (result *ExecuteResult) GetRowsAffected() (int64, error){
	return result.RowsAffected, nil
}