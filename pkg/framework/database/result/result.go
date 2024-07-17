package result

import (
	"encoding/json"
	"fmt"
)

type Row map[string]interface{}

type RowCollection struct {
	Rows []Row
}

func (rc *RowCollection) Add(row Row) {
	rc.Rows = append(rc.Rows, row)
}

func (rc *RowCollection) Unmarshal(result interface{}) error {
	jsonData, err := json.Marshal(rc.Rows)
	if err != nil {
		return fmt.Errorf("error marshaling maps to JSON: %w", err)
	}

	err = json.Unmarshal(jsonData, result)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON to struct: %w", err)
	}

	return nil
}
