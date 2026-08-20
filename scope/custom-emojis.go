package scope

/******************************************
 * Custom Emoji API Methods
 * Each site can define and upload its own custom emoji to be attached to profiles or statuses.
 * https://docs.joinmastodon.org/methods/custom_emojis/
 ******************************************/

// GetCustomEmojis is the OAuth scope required by GET /api/v1/custom_emojis.
// https://docs.joinmastodon.org/methods/custom_emojis/#get
const GetCustomEmojis = Public
