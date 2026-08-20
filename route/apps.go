package route

/******************************************
 * Apps API Methods
 * Register client applications that can be used to obtain OAuth tokens
 * https://docs.joinmastodon.org/methods/apps/
 ******************************************/

// PostApplication is the route for POST /api/v1/apps.
// https://docs.joinmastodon.org/methods/apps/#create
const PostApplication = "/api/v1/apps"

// GetApplication_VerifyCredentials is the route for GET /api/v1/apps/verify_credentials.
// https://docs.joinmastodon.org/methods/apps/#verify_credentials
const GetApplication_VerifyCredentials = "/api/v1/apps/verify_credentials" // #nosec G101 -- a URL path, not a credential
