package auth

import "context"

type (
	Provider interface {
		VerifyIdentityToken(ctx context.Context, jwt string) (*SessionInfo, error)
		GetIdentityToken(ctx context.Context) (*IdentityToken, error)
	}

	//interface only for partner tokens
	PartnerTokenProvider interface {
		GetJWTToken(ctx context.Context) (*JwtToken, error)
	}
)
