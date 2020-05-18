package common

import "fmt"

func GetBaseURL(cc string) string {
	return fmt.Sprintf(BaseUrl, cc)
}

func GetBaseURLFromAuthFunc(f AuthFunc) string {
	params := f("")
	return GetBaseURL(params.Get("clientCode"))
}
