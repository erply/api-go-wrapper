package auth

import (
	"errors"
	"github.com/erply/api-go-wrapper/pkg/common"
	"github.com/stretchr/testify/assert"
	"net/http"
	"sync"
	"testing"
)

func TestGetSessionKeyUserSuccess(t *testing.T) {
	sessKeyUserToGive := SessionKeyUser{
		UserID:             "123",
		UserName:           "someName",
		EmployeeName:       "someEmplName",
		EmployeeID:         "someEmplID",
		GroupID:            "someGroupID",
		GroupName:          "someGroupName",
		IPAddress:          "1.1.1.1",
		SessionKey:         "someSess",
		SessionLength:      10,
		LoginUrl:           "http://LoginUrl.com",
		BerlinPOSVersion:   "123",
		BerlinPOSAssetsURL: "http://BerlinPOSAssetsURL.com",
		EpsiURL:            "http://EpsiURL.com",
		IdentityToken:      "identityToken",
		Token:              "token",
	}
	payload := SessionKeyUserResponse{
		Records: []SessionKeyUser{sessKeyUserToGive},
	}

	bodyMock := common.NewMockFromStruct(payload)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       bodyMock,
	}

	cl := &common.ClientMock{
		ErrToGive:      nil,
		ResponseToGive: resp,
		Lock:           sync.Mutex{},
	}

	sessKeyUserActual, err := GetSessionKeyUser("sess123", "code123", cl)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, sessKeyUserToGive, sessKeyUserActual)
	assert.Len(t, cl.Requests, 1)
	if len(cl.Requests) != 1 {
		return
	}

	req := cl.Requests[0]
	assert.Equal(
		t,
		"https://code123.erply.com/api/?clientCode=code123&doNotGenerateIdentityToken=1&request=getSessionKeyUser&sessionKey=sess123",
		req.URL.String(),
	)
	assert.Equal(t, "application/json", req.Header.Get("Accept"))
	assert.True(t, bodyMock.WasClosed)
}

func TestGetSessionKeyUserInvalidBody(t *testing.T) {
	cl := &common.ClientMock{
		ErrToGive: nil,
		ResponseToGive: &http.Response{
			StatusCode: http.StatusOK,
			Body: &common.BodyMock{
				Body:       common.NewMockFromStr("lala"),
				WasClosed:  false,
				CloseError: nil,
			},
		},
		Lock: sync.Mutex{},
	}

	_, err := GetSessionKeyUser("sess124", "code124", cl)
	assert.Error(t, err)
	if err == nil {
		return
	}
	assert.Contains(t, err.Error(), "ERPLY API: failed to decode SessionKeyUserResponse")
}

func TestGetSessionKeyUserZeroRecords(t *testing.T) {
	payload := SessionKeyUserResponse{
		Records: []SessionKeyUser{},
	}

	bodyMock := common.NewMockFromStruct(payload)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       bodyMock,
	}

	cl := &common.ClientMock{
		ErrToGive:      nil,
		ResponseToGive: resp,
		Lock:           sync.Mutex{},
	}

	_, err := GetSessionKeyUser("sess125", "code125", cl)
	assert.Error(t, err)
	if err == nil {
		return
	}
	assert.Contains(t, err.Error(), "ERPLY API: getSessionKeyUser: no records in response status")
}

func TestGetSessionKeyUserError(t *testing.T) {
	payload := SessionKeyUserResponse{
		Records: []SessionKeyUser{},
	}

	bodyMock := common.NewMockFromStruct(payload)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       bodyMock,
	}

	cl := &common.ClientMock{
		ErrToGive:      errors.New("some bad error"),
		ResponseToGive: resp,
		Lock:           sync.Mutex{},
	}

	_, err := GetSessionKeyUser("sess126", "code126", cl)
	assert.Error(t, err)
	if err == nil {
		return
	}
	assert.EqualError(t, err, "ERPLY API: failed to call getSessionKeyUser request: some bad error status: Error")
}

func TestGetSessionKeyUserWrongRespCode(t *testing.T) {
	sessKeyUserToGive := SessionKeyUser{UserID: "123"}
	payload := SessionKeyUserResponse{
		Records: []SessionKeyUser{sessKeyUserToGive},
	}

	bodyMock := common.NewMockFromStruct(payload)
	resp := &http.Response{
		StatusCode: http.StatusBadRequest,
		Body:       bodyMock,
	}

	cl := &common.ClientMock{
		ErrToGive:      nil,
		ResponseToGive: resp,
		Lock:           sync.Mutex{},
	}

	_, err := GetSessionKeyUser("sess127", "code127", cl)
	assert.Error(t, err)
	if err == nil {
		return
	}
	assert.Contains(t, err.Error(), "wrong response status code: 400")
}

func TestVerifyUser(t *testing.T) {
	const (
		username = ""
		password = ""
		cc       = ""
	)
	sk, err := VerifyUser(username, password, cc, &http.Client{})

	if err != nil {
		t.Error(err)
		return
	}
	if sk == "" {
		t.Error("did not get sk")
		return
	}
}
