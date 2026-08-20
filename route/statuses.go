package route

/******************************************
 * Statuses API Methods
 * Publish, interact, and view information about statuses
 * https://docs.joinmastodon.org/methods/statuses/
 ******************************************/

// PostStatus is the route for POST /api/v1/statuses.
// https://docs.joinmastodon.org/methods/statuses/#create
const PostStatus = "/api/v1/statuses"

// GetStatus is the route for GET /api/v1/statuses/:id.
// https://docs.joinmastodon.org/methods/statuses/#get
const GetStatus = "/api/v1/statuses/:id"

// DeleteStatus is the route for DELETE /api/v1/statuses/:id.
// https://docs.joinmastodon.org/methods/statuses/#delete
const DeleteStatus = "/api/v1/statuses/:id"

// GetStatus_Context is the route for GET /api/v1/statuses/:id/context.
// https://docs.joinmastodon.org/methods/statuses/#context
const GetStatus_Context = "/api/v1/statuses/:id/context"

// PostStatus_Translate is the route for POST /api/v1/statuses/:id/translate.
// https://docs.joinmastodon.org/methods/statuses/#translate
const PostStatus_Translate = "/api/v1/statuses/:id/translate"

// GetStatus_RebloggedBy is the route for GET /api/v1/statuses/:id/reblogged_by.
// https://docs.joinmastodon.org/methods/statuses/#reblogged_by
const GetStatus_RebloggedBy = "/api/v1/statuses/:id/reblogged_by"

// GetStatus_FavouritedBy is the route for GET /api/v1/statuses/:id/favourited_by.
// https://docs.joinmastodon.org/methods/statuses/#favourited_by
const GetStatus_FavouritedBy = "/api/v1/statuses/:id/favourited_by"

// PostStatus_Favourite is the route for POST /api/v1/statuses/:id/favourite.
// https://docs.joinmastodon.org/methods/statuses/#favourite
const PostStatus_Favourite = "/api/v1/statuses/:id/favourite"

// PostStatus_Unfavourite is the route for POST /api/v1/statuses/:id/unfavourite.
// https://docs.joinmastodon.org/methods/statuses/#unfavourite
const PostStatus_Unfavourite = "/api/v1/statuses/:id/unfavourite"

// PostStatus_Reblog is the route for POST /api/v1/statuses/:id/reblog.
// https://docs.joinmastodon.org/methods/statuses/#reblog
const PostStatus_Reblog = "/api/v1/statuses/:id/reblog"

// PostStatus_Unreblog is the route for POST /api/v1/statuses/:id/unreblog.
// https://docs.joinmastodon.org/methods/statuses/#unreblog
const PostStatus_Unreblog = "/api/v1/statuses/:id/unreblog"

// PostStatus_Bookmark is the route for POST /api/v1/statuses/:id/bookmark.
// https://docs.joinmastodon.org/methods/statuses/#bookmark
const PostStatus_Bookmark = "/api/v1/statuses/:id/bookmark"

// PostStatus_Unbookmark is the route for POST /api/v1/statuses/:id/unbookmark.
// https://docs.joinmastodon.org/methods/statuses/#unbookmark
const PostStatus_Unbookmark = "/api/v1/statuses/:id/unbookmark"

// PostStatus_Mute is the route for POST /api/v1/statuses/:id/mute.
// https://docs.joinmastodon.org/methods/statuses/#mute
const PostStatus_Mute = "/api/v1/statuses/:id/mute"

// PostStatus_Unmute is the route for POST /api/v1/statuses/:id/unmute.
// https://docs.joinmastodon.org/methods/statuses/#unmute
const PostStatus_Unmute = "/api/v1/statuses/:id/unmute"

// PostStatus_Pin is the route for POST /api/v1/statuses/:id/pin.
// https://docs.joinmastodon.org/methods/statuses/#pin
const PostStatus_Pin = "/api/v1/statuses/:id/pin"

// PostStatus_Unpin is the route for POST /api/v1/statuses/:id/unpin.
// https://docs.joinmastodon.org/methods/statuses/#unpin
const PostStatus_Unpin = "/api/v1/statuses/:id/unpin"

// PutStatus is the route for PUT /api/v1/statuses/:id.
// https://docs.joinmastodon.org/methods/statuses/#edit
const PutStatus = "/api/v1/statuses/:id"

// GetStatus_History is the route for GET /api/v1/statuses/:id/history.
// https://docs.joinmastodon.org/methods/statuses/#history
const GetStatus_History = "/api/v1/statuses/:id/history"

// GetStatus_Source is the route for GET /api/v1/statuses/:id/source.
// https://docs.joinmastodon.org/methods/statuses/#source
const GetStatus_Source = "/api/v1/statuses/:id/source"
