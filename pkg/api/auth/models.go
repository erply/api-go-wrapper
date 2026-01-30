package auth

import (
	"encoding/json"
	common2 "github.com/erply/api-go-wrapper/pkg/api/common"
)

type (
	VerifyUserResponse struct {
		Status  common2.Status   `json:"status"`
		Records []SessionKeyUser `json:"records"`
	}
	SwitchUserResponse struct {
		Status  common2.Status   `json:"status"`
		Records []SessionKeyUser `json:"records"`
	}

	verifyIdentityTokenResponse struct {
		Status common2.Status `json:"status"`
		Result SessionInfo    `json:"records"`
	}

	SessionInfo struct {
		SessionKey string `json:"sessionKey"`
	}

	getIdentityTokenResponse struct {
		Status common2.Status `json:"status"`
		Result IdentityToken  `json:"records"`
	}
	IdentityToken struct {
		Jwt string `json:"identityToken"`
	}
	JwtTokenResponse struct {
		Status  common2.Status `json:"status"`
		Records JwtToken       `json:"records"`
	}
	JwtTokenSafeResponse struct {
		Status  common2.Status  `json:"status"`
		Records json.RawMessage `json:"records"`
	}
	JwtToken struct {
		Token string `json:"token"`
	}

	SessionKeyUserResponse struct {
		Records []SessionKeyUser `json:"records"`
	}

	SessionKeyUser struct {
		UserID               string `json:"userID"`
		UserName             string `json:"userName"`
		EmployeeName         string `json:"employeeName"`
		EmployeeID           string `json:"employeeID"`
		GroupID              string `json:"groupID"`
		GroupName            string `json:"groupName"`
		IPAddress            string `json:"ipAddress"`
		SessionKey           string `json:"sessionKey"`
		SessionLength        int    `json:"sessionLength"`
		LoginUrl             string `json:"loginUrl"`
		BerlinPOSVersion     string `json:"berlinPOSVersion"`
		BerlinPOSAssetsURL   string `json:"berlinPOSAssetsURL"`
		EpsiURL              string `json:"epsiURL"`
		IdentityToken        string `json:"identityToken"`
		Token                string `json:"token"`
		CustomerRegistryURLs []struct {
			Priority int64  `json:"priority"`
			Token    string `json:"token"`
			URL      string `json:"url"`
			Weight   int64  `json:"weight"`
		} `json:"customerRegistryURLs"`
	}

	SessionKeyInfoResponse struct {
		Status  common2.Status   `json:"status"`
		Records []SessionKeyInfo `json:"records"`
	}
	SessionKeyInfo struct {
		CreationUnixTime string `json:"creationUnixTime"`
		ExpireUnixTime   string `json:"expireUnixTime"`
	}
)
