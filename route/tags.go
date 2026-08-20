package route

/******************************************
 * Tags API Methods
 * View information about or follow/unfollow hashtags
 * https://docs.joinmastodon.org/methods/tags/
 ******************************************/

// GetTag is the route for GET /api/v1/tags/:id.
// https://docs.joinmastodon.org/methods/tags/#get
const GetTag = "/api/v1/tags/:id"

// PostTag_Follow is the route for POST /api/v1/tags/:id/follow.
// https://docs.joinmastodon.org/methods/tags/#follow
const PostTag_Follow = "/api/v1/tags/:id/follow"

// PostTag_Unfollow is the route for POST /api/v1/tags/:id/unfollow.
// https://docs.joinmastodon.org/methods/tags/#unfollow
const PostTag_Unfollow = "/api/v1/tags/:id/unfollow"
