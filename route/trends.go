package route

/******************************************
 * Trends API Methods
 * View hashtags that are currently being used more frequently than usual.
 * https://docs.joinmastodon.org/methods/trends/
 ******************************************/

// GetTrends is the route for GET /api/v1/trends.
// https://docs.joinmastodon.org/methods/trends/#tags
const GetTrends = "/api/v1/trends"

// GetTrends_Statuses is the route for GET /api/v1/trends/statuses.
// https://docs.joinmastodon.org/methods/trends/#statuses
const GetTrends_Statuses = "/api/v1/trends/statuses"

// GetTrends_Links is the route for GET /api/v1/trends/links.
// https://docs.joinmastodon.org/methods/trends/#links
const GetTrends_Links = "/api/v1/trends/links"
