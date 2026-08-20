package route

/******************************************
 * OAuth API Methods
 * Generate and manage OAuth tokens
 * https://docs.joinmastodon.org/methods/oauth/
******************************************/

// GetOAuth_Authorize is the route for GET /oauth/authorize.
// https://docs.joinmastodon.org/methods/oauth/#authorize
const GetOAuth_Authorize = "/oauth/authorize"

// PostOAuth_Token is the route for POST /oauth/token.
// https://docs.joinmastodon.org/methods/oauth/#token
const PostOAuth_Token = "/oauth/token" // #nosec G101 -- a URL path, not a credential

// PostOAuth_Revoke is the route for POST /oauth/revoke.
// https://docs.joinmastodon.org/methods/oauth/#revoke
const PostOAuth_Revoke = "/oauth/revoke"
