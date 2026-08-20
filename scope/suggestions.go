package scope

/******************************************
 * Suggestions API Methods
 * Server-generated suggestions on who to follow, based on previous
 * positive interactions
 * https://docs.joinmastodon.org/methods/suggestions/
 ******************************************/

// GetSuggestions is the OAuth scope required by GET /api/v2/suggestions.
// https://docs.joinmastodon.org/methods/suggestions/#v2
const GetSuggestions = Read

// DeleteSuggestion is the OAuth scope required by DELETE /api/v1/suggestions/:account_id.
// https://docs.joinmastodon.org/methods/suggestions/#remove
const DeleteSuggestion = Read
