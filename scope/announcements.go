package scope

/******************************************
 * Announcements API Methods
 * For announcements set by administration
 * https://docs.joinmastodon.org/methods/announcements/
 ******************************************/

// GetAnnouncements is the OAuth scope required by GET /api/v1/announcements.
// https://docs.joinmastodon.org/methods/announcements/#get
const GetAnnouncements = Private

// PostAnnouncement_Dismiss is the OAuth scope required by POST /api/v1/announcements/:id/dismiss.
// https://docs.joinmastodon.org/methods/announcements/#dismiss
const PostAnnouncement_Dismiss = WriteAccounts

// PutAnnouncement_Reaction is the OAuth scope required by PUT /api/v1/announcements/:id/reactions/:name.
// https://docs.joinmastodon.org/methods/announcements/#put-reactions
const PutAnnouncement_Reaction = WriteFavourites

// DeleteAnnouncement_Reaction is the OAuth scope required by DELETE /api/v1/announcements/:id/reactions/:name.
// https://docs.joinmastodon.org/methods/announcements/#delete-reactions
const DeleteAnnouncement_Reaction = WriteFavourites
