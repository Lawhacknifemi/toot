package txn

/******************************************
 * Trends API Methods
 * View hashtags that are currently being used more frequently than usual.
 * https://docs.joinmastodon.org/methods/trends/
 ******************************************/

// GetTrends is the input for GET /api/v1/trends, which returns []Tag.
// https://docs.joinmastodon.org/methods/trends/#tags
type GetTrends struct {
	Host   string `header:"Host"`
	Limit  int    `query:"limit"`  // Maximum number of results to return. Defaults to 10 tags. Max 20 tags.
	Offset int    `query:"offset"` // Skip the first n results.
}

// GetTrends_Statuses is the input for GET /api/v1/trends/statuses, which returns []Status.
// https://docs.joinmastodon.org/methods/trends/#statuses
type GetTrends_Statuses struct {
	Host   string `header:"Host"`
	Limit  int    `query:"limit"`  // Maximum number of results to return. Defaults to 10 tags. Max 20 tags.
	Offset int    `query:"offset"` // Skip the first n results.
}

// GetTrends_Links is the input for GET /api/v1/trends/links, which returns []Link.
// https://docs.joinmastodon.org/methods/trends/#links
type GetTrends_Links struct {
	Host   string `header:"Host"`
	Limit  int    `query:"limit"`  // Maximum number of results to return. Defaults to 10 tags. Max 20 tags.
	Offset int    `query:"offset"` // Skip the first n results.
}
