package auth

import "github.com/erply/api-go-wrapper/pkg/common"

type (
	VerifyUserResponse struct {
		Records []Records `json:"records"`
	}

	Records struct {
		SessionKey string `json:"sessionKey"`
	}
	verifyIdentityTokenResponse struct {
		Status common.Status `json:"status"`
		Result SessionInfo   `json:"records"`
	}

	SessionInfo struct {
		SessionKey string `json:"sessionKey"`
	}

	getIdentityTokenResponse struct {
		Status common.Status `json:"status"`
		Result IdentityToken `json:"records"`
	}
	IdentityToken struct {
		Jwt string `json:"identityToken"`
	}
	JwtTokenResponse struct {
		Status  common.Status `json:"status"`
		Records JwtToken      `json:"records"`
	}
	JwtToken struct {
		Token string `json:"token"`
	}

	SessionKeyUserResponse struct {
		Records []SessionKeyUser `json:"records"`
	}

	SessionKeyUser struct {
		UserID             string `json:"userID"`
		UserName           string `json:"userName"`
		EmployeeName       string `json:"employeeName"`
		EmployeeID         string `json:"employeeID"`
		GroupID            string `json:"groupID"`
		GroupName          string `json:"groupName"`
		IPAddress          string `json:"ipAddress"`
		SessionKey         string `json:"sessionKey"`
		SessionLength      int    `json:"sessionLength"`
		LoginUrl           string `json:"loginUrl"`
		BerlinPOSVersion   string `json:"berlinPOSVersion"`
		BerlinPOSAssetsURL string `json:"berlinPOSAssetsURL"`
		EpsiURL            string `json:"epsiURL"`
		IdentityToken      string `json:"identityToken"`
		Token              string `json:"token"`
	}
)
