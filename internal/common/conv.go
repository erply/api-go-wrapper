package common

import (
	"encoding/json"
)

func ConvertStructToMap(input interface{}) (map[string]interface{}, error) {
	jsonStr, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	rawMap := map[string]interface{}{}
	err = json.Unmarshal(jsonStr, &rawMap)
	if err != nil {
		return nil, err
	}

	return rawMap, nil
}
