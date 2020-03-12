package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GenerateToken(partnerKey string, timestamp int64, request, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(fmt.Sprintf("%s%s%d", partnerKey, request, timestamp)))
	return hex.EncodeToString(h.Sum(nil))

}
