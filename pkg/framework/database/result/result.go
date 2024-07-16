package result

type Row map[string]interface{}

type RowCollection struct {
    Rows []Row
}

func (rc *RowCollection) Add(row Row) {
    rc.Rows = append(rc.Rows, row)
}

func (rc *RowCollection) GetRows() []Row {
    return rc.Rows
}