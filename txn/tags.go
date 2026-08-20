package txn

/******************************************
 * Tags API Methods
 * View information about or follow/unfollow hashtags
 * https://docs.joinmastodon.org/methods/tags/
 ******************************************/

// GetTag is the input for GET /api/v1/tags/:id, which returns Tag.
// https://docs.joinmastodon.org/methods/tags/#get
type GetTag struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostTag_Follow is the input for POST /api/v1/tags/:id/follow, which returns Tag.
// https://docs.joinmastodon.org/methods/tags/#follow
type PostTag_Follow struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostTag_Unfollow is the input for POST /api/v1/tags/:id/unfollow, which returns Tag.
// https://docs.joinmastodon.org/methods/tags/#unfollow
type PostTag_Unfollow struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}
