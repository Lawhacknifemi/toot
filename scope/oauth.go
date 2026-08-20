package scope

/******************************************
 * OAuth API Methods
 * Generate and manage OAuth tokens
 * https://docs.joinmastodon.org/methods/oauth/
******************************************/

// GetOAuth_Authorize is the OAuth scope required by GET /oauth/authorize.
// https://docs.joinmastodon.org/methods/oauth/#authorize
const GetOAuth_Authorize = Public

// PostOAuth_Token is the OAuth scope required by POST /oauth/token.
// https://docs.joinmastodon.org/methods/oauth/#token
const PostOAuth_Token = Public

// PostOAuth_Revoke is the OAuth scope required by POST /oauth/revoke.
// https://docs.joinmastodon.org/methods/oauth/#revoke
const PostOAuth_Revoke = Public
