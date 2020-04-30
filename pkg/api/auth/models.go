package auth

import "github.com/erply/api-go-wrapper/pkg/common"

type (
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
)
