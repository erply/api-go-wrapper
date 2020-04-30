package api

/*
func NewPartnerClient(sessionKey, clientCode, partnerKey string, customCli *http.Client) (IPartnerClient, error) {
	if sessionKey == "" || clientCode == "" || partnerKey == "" {
		return nil, errors.New("sessionKey, clientCode and partnerKey are required")
	}

	cli := erplyClient{
		sessionKey: sessionKey,
		clientCode: clientCode,
		partnerKey: partnerKey,
		url:        fmt.Sprintf(baseURL, clientCode),
		httpClient: getDefaultHTTPClient(),
	}
	if customCli != nil {
		cli.httpClient = customCli
	}
	return &cli, nil
}
*/
