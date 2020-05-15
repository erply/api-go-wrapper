package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

//ClientMock mocks HttpClient interface
type ClientMock struct {
	ErrToGive      error
	ResponseToGive *http.Response
	Requests       []*http.Request
	Lock           sync.Mutex
}

//Do HttpClient interface implementation
func (cm *ClientMock) Do(req *http.Request) (*http.Response, error) {
	cm.Lock.Lock()
	defer cm.Lock.Unlock()

	cm.Requests = append(cm.Requests, req)

	return cm.ResponseToGive, cm.ErrToGive
}

//BodyMock implements resp.Body
type BodyMock struct {
	Body       io.Reader
	WasClosed  bool
	CloseError error
}

//NewFromStr creates BodyMock from a string
func NewMockFromStr(body string) *BodyMock {
	return &BodyMock{
		Body:       strings.NewReader(body),
		WasClosed:  false,
		CloseError: nil,
	}
}

//NewFromStruct creates BodyMock from a struct converted to json or string
func NewMockFromStruct(input interface{}) *BodyMock {
	c, err := json.Marshal(input)
	if err != nil {
		return &BodyMock{
			Body:       strings.NewReader(fmt.Sprintf("%+v", input)),
			WasClosed:  false,
			CloseError: nil,
		}
	}
	buf := bytes.NewBuffer(c)
	return &BodyMock{
		Body:       buf,
		WasClosed:  false,
		CloseError: nil,
	}
}

//Read io.Reader implementation
func (bm *BodyMock) Read(p []byte) (n int, err error) {
	return bm.Body.Read(p)
}

//Close io.Closer implementation
func (bm *BodyMock) Close() error {
	bm.WasClosed = true
	return bm.CloseError
}
