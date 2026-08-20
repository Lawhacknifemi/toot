package txn

/******************************************
 * Custom Emoji API Methods
 * Each site can define and upload its own custom emoji to be attached to profiles or statuses.
 * https://docs.joinmastodon.org/methods/custom_emojis/
 ******************************************/

// GetCustomEmojis is the input for GET /api/v1/custom_emojis, which returns []CustomEmoji.
// https://docs.joinmastodon.org/methods/custom_emojis/#get
type GetCustomEmojis struct{}
