package txn

/******************************************
 * Apps API Methods
 * Register client applications that can be used to obtain OAuth tokens
 * https://docs.joinmastodon.org/methods/apps/
 ******************************************/

// PostApplication is the input for POST /api/v1/apps, which returns Application.
// https://docs.joinmastodon.org/methods/apps/#create
type PostApplication struct {
	Host         string `header:"Host"`
	ClientName   string `form:"client_name"   json:"client_name"`
	RedirectURIs string `form:"redirect_uris" json:"redirect_uris"`
	Scopes       string `form:"scopes"        json:"scopes"`
	Website      string `form:"website"       json:"website"`
}

// GetApplication_VerifyCredentials is the input for GET /api/v1/apps/verify_credentials, which returns Application.
// https://docs.joinmastodon.org/methods/apps/#verify_credentials
type GetApplication_VerifyCredentials struct {
	Host string `header:"Host"`
}
