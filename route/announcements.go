package route

/******************************************
 * Announcements API Methods
 * For announcements set by administration
 * https://docs.joinmastodon.org/methods/announcements/
 ******************************************/

// GetAnnouncements is the route for GET /api/v1/announcements.
// https://docs.joinmastodon.org/methods/announcements/#get
const GetAnnouncements = "/api/v1/announcements"

// PostAnnouncement_Dismiss is the route for POST /api/v1/announcements/:id/dismiss.
// https://docs.joinmastodon.org/methods/announcements/#dismiss
const PostAnnouncement_Dismiss = "/api/v1/announcements/:id/dismiss"

// PutAnnouncement_Reaction is the route for PUT /api/v1/announcements/:id/reactions/:name.
// https://docs.joinmastodon.org/methods/announcements/#put-reactions
const PutAnnouncement_Reaction = "/api/v1/announcements/:id/reactions/:name"

// DeleteAnnouncement_Reaction is the route for DELETE /api/v1/announcements/:id/reactions/:name.
// https://docs.joinmastodon.org/methods/announcements/#delete-reactions
const DeleteAnnouncement_Reaction = "/api/v1/announcements/:id/reactions/:name"
