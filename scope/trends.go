package scope

/******************************************
 * Trends API Methods
 * View hashtags that are currently being used more frequently than usual.
 * https://docs.joinmastodon.org/methods/trends/
 ******************************************/

// GetTrends is the OAuth scope required by GET /api/v1/trends.
// https://docs.joinmastodon.org/methods/trends/#tags
const GetTrends = Public

// GetTrends_Statuses is the OAuth scope required by GET /api/v1/trends/statuses.
// https://docs.joinmastodon.org/methods/trends/#statuses
const GetTrends_Statuses = Public

// GetTrends_Links is the OAuth scope required by GET /api/v1/trends/links.
// https://docs.joinmastodon.org/methods/trends/#links
const GetTrends_Links = Public
