package servicediscovery

import "context"

type ServiceDiscoverer interface {
	GetServiceEndpoints(ctx context.Context) (*ServiceEndpoints, error)
}
