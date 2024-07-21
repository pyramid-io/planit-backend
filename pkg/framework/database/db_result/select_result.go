package db_result

import (
	"encoding/json"
	"fmt"
)

type SelectResult struct {
	RowsCollection []Row
}
type Row map[string]interface{}

func (sr *SelectResult) Add(row Row) {
	sr.RowsCollection = append(sr.RowsCollection, row)
}

func (sr *SelectResult) Unmarshal(result interface{}) error {
	jsonData, err := json.Marshal(sr.RowsCollection)
	if err != nil {
		return fmt.Errorf("error marshaling maps to JSON: %w", err)
	}

	err = json.Unmarshal(jsonData, result)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON to struct: %w", err)
	}

	return nil
}
