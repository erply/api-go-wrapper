package common

import (
	"fmt"
)

var activeBaseDomain = BaseDomain

func SetBaseDomain(domain string) {
	activeBaseDomain = domain
}

func GetBaseURL(cc string) string {
	return fmt.Sprintf(BaseUrl, cc, activeBaseDomain)
}

func GetBaseURLFromAuthFunc(f AuthFunc) string {
	params := f("")
	return GetBaseURL(params.Get("clientCode"))
}
