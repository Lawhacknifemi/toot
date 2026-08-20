package scope

/******************************************
 * Statuses API Methods
 * Publish, interact, and view information about statuses
 * https://docs.joinmastodon.org/methods/statuses/
 ******************************************/

// PostStatus is the OAuth scope required by POST /api/v1/statuses.
// https://docs.joinmastodon.org/methods/statuses/#create
const PostStatus = WriteStatuses

// GetStatus is the OAuth scope required by GET /api/v1/statuses/:id.
// https://docs.joinmastodon.org/methods/statuses/#get
const GetStatus = ReadStatuses

// DeleteStatus is the OAuth scope required by DELETE /api/v1/statuses/:id.
// https://docs.joinmastodon.org/methods/statuses/#delete
const DeleteStatus = WriteStatuses

// GetStatus_Context is the OAuth scope required by GET /api/v1/statuses/:id/context.
// https://docs.joinmastodon.org/methods/statuses/#context
const GetStatus_Context = ReadStatuses

// PostStatus_Translate is the OAuth scope required by POST /api/v1/statuses/:id/translate.
// https://docs.joinmastodon.org/methods/statuses/#translate
const PostStatus_Translate = WriteStatuses

// GetStatus_RebloggedBy is the OAuth scope required by GET /api/v1/statuses/:id/reblogged_by.
// https://docs.joinmastodon.org/methods/statuses/#reblogged_by
const GetStatus_RebloggedBy = ReadStatuses

// GetStatus_FavouritedBy is the OAuth scope required by GET /api/v1/statuses/:id/favourited_by.
// https://docs.joinmastodon.org/methods/statuses/#favourited_by
const GetStatus_FavouritedBy = ReadStatuses

// PostStatus_Favourite is the OAuth scope required by POST /api/v1/statuses/:id/favourite.
// https://docs.joinmastodon.org/methods/statuses/#favourite
const PostStatus_Favourite = WriteStatuses

// PostStatus_Unfavourite is the OAuth scope required by POST /api/v1/statuses/:id/unfavourite.
// https://docs.joinmastodon.org/methods/statuses/#unfavourite
const PostStatus_Unfavourite = WriteStatuses

// PostStatus_Reblog is the OAuth scope required by POST /api/v1/statuses/:id/reblog.
// https://docs.joinmastodon.org/methods/statuses/#reblog
const PostStatus_Reblog = WriteStatuses

// PostStatus_Unreblog is the OAuth scope required by POST /api/v1/statuses/:id/unreblog.
// https://docs.joinmastodon.org/methods/statuses/#unreblog
const PostStatus_Unreblog = WriteStatuses

// PostStatus_Bookmark is the OAuth scope required by POST /api/v1/statuses/:id/bookmark.
// https://docs.joinmastodon.org/methods/statuses/#bookmark
const PostStatus_Bookmark = WriteStatuses

// PostStatus_Unbookmark is the OAuth scope required by POST /api/v1/statuses/:id/unbookmark.
// https://docs.joinmastodon.org/methods/statuses/#unbookmark
const PostStatus_Unbookmark = WriteStatuses

// PostStatus_Mute is the OAuth scope required by POST /api/v1/statuses/:id/mute.
// https://docs.joinmastodon.org/methods/statuses/#mute
const PostStatus_Mute = WriteStatuses

// PostStatus_Unmute is the OAuth scope required by POST /api/v1/statuses/:id/unmute.
// https://docs.joinmastodon.org/methods/statuses/#unmute
const PostStatus_Unmute = WriteStatuses

// PostStatus_Pin is the OAuth scope required by POST /api/v1/statuses/:id/pin.
// https://docs.joinmastodon.org/methods/statuses/#pin
const PostStatus_Pin = WriteStatuses

// PostStatus_Unpin is the OAuth scope required by POST /api/v1/statuses/:id/unpin.
// https://docs.joinmastodon.org/methods/statuses/#unpin
const PostStatus_Unpin = WriteStatuses

// PutStatus is the OAuth scope required by PUT /api/v1/statuses/:id.
// https://docs.joinmastodon.org/methods/statuses/#edit
const PutStatus = WriteStatuses

// GetStatus_History is the OAuth scope required by GET /api/v1/statuses/:id/history.
// https://docs.joinmastodon.org/methods/statuses/#history
const GetStatus_History = ReadStatuses

// GetStatus_Source is the OAuth scope required by GET /api/v1/statuses/:id/source.
// https://docs.joinmastodon.org/methods/statuses/#source
const GetStatus_Source = ReadStatuses
