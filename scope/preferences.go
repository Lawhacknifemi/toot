package scope

/******************************************
* Preferences API Methods
* Preferred common behaviors to be shared across clients.
* https://docs.joinmastodon.org/methods/preferences/
******************************************/

// GetPreferences is the OAuth scope required by GET /api/v1/preferences.
// https://docs.joinmastodon.org/methods/preferences/#get
const GetPreferences = ReadAccounts
