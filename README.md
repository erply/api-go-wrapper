## ERPLY API Go SDK

For available requests please check `pkg/api/IClient.go` file

### Example usage as a service

```go
import (
	"github.com/pkg/errors"
	"github.com/erply/api-go-wrapper/pkg/api"
	"strconv"
	"strings"
)

type erplyApiService struct {
	api.IClient
}

func NewErplyApiService(sessionKey, clientCode string) *erplyApiService {
	return &erplyApiService{api.NewClient(sessionKey, clientCode, nil)}
}

//GetSalesDocument erply API request
func (s *erplyApiService) getPointsOfSale(posID string) (string, error) {
	res, err := s.GetPointsOfSaleByID(posID)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(res.WarehouseID), nil
}

//verifyIdentityToken erply API request
func (s *erplyApiService) verifyIdentityToken(jwt string) (string, error) {
	res, err := s.VerifyIdentityToken(jwt)
	if err != nil {
		if strings.Contains(err.Error(), "1000") {
			return "", errors.New("jwt expired")
		}
	}
	return res.SessionKey, nil
}

//getIdentityToken erply API request
func (s *erplyApiService) getIdentityToken() (string, error) {
	res, err := s.GetIdentityToken()
	if err != nil {
		if strings.Contains(err.Error(), "1054") {
			return "", errors.New("API session key expired")
		}
	}
	return res.Jwt, nil
}
```