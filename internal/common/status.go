package common

type Status struct {
	Request           string  `json:"request"`
	RequestUnixTime   int     `json:"requestUnixTime"`
	ResponseStatus    string  `json:"responseStatus"`
	ErrorCode         int     `json:"errorCode"`
	ErrorField        string  `json:"errorField"`
	GenerationTime    float64 `json:"generationTime"`
	RecordsTotal      int     `json:"recordsTotal"`
	RecordsInResponse int     `json:"recordsInResponse"`
}
