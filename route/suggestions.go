package route

/******************************************
 * Suggestions API Methods
 * Server-generated suggestions on who to follow, based on previous
 * positive interactions
 * https://docs.joinmastodon.org/methods/suggestions/
 ******************************************/

// GetSuggestions is the route for GET /api/v2/suggestions.
// https://docs.joinmastodon.org/methods/suggestions/#v2
const GetSuggestions = "/api/v2/suggestions"

// DeleteSuggestion is the route for DELETE /api/v1/suggestions/:account_id.
// https://docs.joinmastodon.org/methods/suggestions/#remove
const DeleteSuggestion = "/api/v1/suggestions/:account_id"
