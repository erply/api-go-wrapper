package examples

import (
	"context"
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/api"
)

func main() {
	const (
		sk         = ""
		cc         = ""
		partnerKey = ""
	)

	cli, err := api.NewClient(sk, cc, nil)
	if err != nil {
		panic(err)
	}

	endpoints, err := cli.ServiceDiscoverer.GetServiceEndpoints(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(endpoints)

	partnerCli, err := api.NewPartnerClient(sk, cc, partnerKey, nil)
	jwt, err := partnerCli.PartnerTokenProvider.GetJWTToken(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(jwt)
}
