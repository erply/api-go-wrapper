package common

import (
	"encoding/json"
	"net/http"
)

func ExtractBulkFiltersFromRequest(r *http.Request) (res map[string]interface{}, err error) {
	err = r.ParseForm()
	if err != nil {
		return
	}

	res = make(map[string]interface{})
	for key, vals := range r.Form {
		res[key] = vals[0]
	}

	var requests []map[string]interface{}
	requestsRaw, ok := res["requests"]
	if ok {
		err = json.Unmarshal([]byte(requestsRaw.(string)), &requests)
		if err != nil {
			return
		}
	}
	res["requests"] = requests
	return
}
