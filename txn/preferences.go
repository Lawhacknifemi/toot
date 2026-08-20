package txn

/******************************************
* Preferences API Methods
* Preferred common behaviors to be shared across clients.
* https://docs.joinmastodon.org/methods/preferences/
******************************************/

// GetPreferences is the input for GET /api/v1/preferences, which returns Preferences.
// https://docs.joinmastodon.org/methods/preferences/#get
type GetPreferences struct {
	Host string `header:"Host"`
}
