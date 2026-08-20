package txn

/******************************************
 * Suggestions API Methods
 * Server-generated suggestions on who to follow, based on previous
 * positive interactions
 * https://docs.joinmastodon.org/methods/suggestions/
 ******************************************/

// GetSuggestions is the input for GET /api/v2/suggestions, which returns []Suggestion.
// https://docs.joinmastodon.org/methods/suggestions/#v2
type GetSuggestions struct {
	Host  string `header:"Host"`
	Limit int    `query:"limit"`
}

// DeleteSuggestion is the input for DELETE /api/v1/suggestions/:account_id.
// https://docs.joinmastodon.org/methods/suggestions/#remove
type DeleteSuggestion struct {
	Host      string `header:"Host"`
	AccountID string `param:"account_id"`
}
