package object

// https://docs.joinmastodon.org/entities/Token/
type Token struct {
	AccessToken  string `json:"access_token"`            // An OAuth token to be used for authorization.
	TokenType    string `json:"token_type"`              // The OAuth token type. Mastodon uses Bearer tokens.
	Scope        string `json:"scope"`                   // The OAuth scopes granted by this token, space-separated.
	CreatedAt    int64  `json:"created_at"`              // When the token was generated. (Unix Timestamp)
	ExpiresIn    int64  `json:"expires_in,omitempty"`    // Lifetime in seconds of the access token (RFC 6749 §5.1). Omitted when the token does not expire.
	RefreshToken string `json:"refresh_token,omitempty"` // The refresh token used to obtain a new access token (RFC 6749 §6). Omitted when no refresh token is issued.
}
