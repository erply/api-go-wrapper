package common

import "github.com/erply/api-go-wrapper/pkg/api/errors"

type Status struct {
	Request           string          `json:"request"`
	RequestUnixTime   int             `json:"requestUnixTime"`
	ResponseStatus    string          `json:"responseStatus"`
	ErrorCode         errors.ApiError `json:"errorCode"`
	ErrorField        string          `json:"errorField"`
	GenerationTime    float64         `json:"generationTime"`
	RecordsTotal      int             `json:"recordsTotal"`
	RecordsInResponse int             `json:"recordsInResponse"`
}

type StatusBulk struct {
	RequestName string `json:"requestName"`
	RequestID   string `json:"requestID"`
	Status
}
