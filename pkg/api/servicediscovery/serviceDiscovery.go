package servicediscovery

/*
type ServiceDiscoverer interface {
	GetServiceEndpoints(ctx context.Context) (*ServiceEndpoints, error)
}

type getServiceEndpointsResponse struct {
	Status  common.Status
	Records []ServiceEndpoints `json:"records"`
}

type ServiceEndpoints struct {
	Cafa        Endpoint `json:"cafa"`
	Pim         Endpoint `json:"pim"`
	Wms         Endpoint `json:"wms"`
	Promotion   Endpoint `json:"promotion"`
	Reports     Endpoint `json:"reports"`
	Json        Endpoint `json:"json"`
	Assignments Endpoint `json:"assignments"`
}
type Endpoint struct {
	Url           string `json:"url"`
	Documentation string `json:"documentation"`
}

func (cli *Client) GetServiceEndpoints(ctx context.Context) (*ServiceEndpoints, error) {
	const method = "getServiceEndpoints"
	resp, err := cli.SendRequest(ctx, method, map[string]string{})
	if err != nil {
		return nil, err
	}

	res := &getServiceEndpointsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to decode %s response", method))
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	if len(res.Records) < 1 {
		return nil, errors.New("no records in response")
	}
	return &res.Records[0], nil
}
*/
