package scope

/******************************************
 * Apps API Methods
 * Register client applications that can be used to obtain OAuth tokens
 * https://docs.joinmastodon.org/methods/apps/
 ******************************************/

// PostApplication is the OAuth scope required by POST /api/v1/apps.
// https://docs.joinmastodon.org/methods/apps/#create
const PostApplication = Public

// GetApplication_VerifyCredentials is the OAuth scope required by GET /api/v1/apps/verify_credentials.
// https://docs.joinmastodon.org/methods/apps/#verify_credentials
const GetApplication_VerifyCredentials = Private
