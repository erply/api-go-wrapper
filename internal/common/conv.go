package common

import (
	"encoding/json"
	"strings"
)

func ConvertStructToMap(input interface{}) (map[string]string, error) {
	supplierJson, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	rawMap := map[string]json.RawMessage{}
	err = json.Unmarshal(supplierJson, &rawMap)
	if err != nil {
		return nil, err
	}

	bulkFilterMap := make(map[string]string, len(rawMap))
	for key, val := range rawMap {
		bulkFilterMap[key] = strings.Trim(string(val), `"`)
	}

	return bulkFilterMap, nil
}
