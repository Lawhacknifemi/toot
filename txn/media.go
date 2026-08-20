package txn

/******************************************
 * Media API Methods
 * Attach media to authored statuses. See Using Mastodon > Posting toots > Attachments
 * for more information about size and format limits
 * https://docs.joinmastodon.org/methods/media/
 ******************************************/

// PostMedia is the input for POST /api/v2/media, which returns MediaAttachment.
// https://docs.joinmastodon.org/methods/media/#v2
type PostMedia struct {
	Host        string `header:"Host"`
	File        string `form:"file"`
	Thumbnail   string `form:"thumbnail"`
	Description string `form:"description"`
	Focus       string `form:"focus"`
}
