package scope

/******************************************
 * Mutes API Methods
 * View your mutes. See also /accounts/:id/{mute,unmute}
 * https://docs.joinmastodon.org/methods/mutes/
 ******************************************/

// GetMutes is the OAuth scope required by GET /api/v1/mutes.
// https://docs.joinmastodon.org/methods/mutes/#get
const GetMutes = ReadMutes
