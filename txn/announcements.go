package txn

/******************************************
 * Announcements API Methods
 * For announcements set by administration
 * https://docs.joinmastodon.org/methods/announcements/
 ******************************************/

// GetAnnouncements is the input for GET /api/v1/announcements, which returns []Announcement.
// https://docs.joinmastodon.org/methods/announcements/#get
type GetAnnouncements struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	WithDismissed bool   `query:"with_dismissed"`
}

// PostAnnouncement_Dismiss is the input for POST /api/v1/announcements/:id/dismiss, which returns an empty object.
// https://docs.joinmastodon.org/methods/announcements/#dismiss
type PostAnnouncement_Dismiss struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	ID            string `param:"id"`
}

// PutAnnouncement_Reaction is the input for PUT /api/v1/announcements/:id/reactions/:name, which returns an empty object.
// https://docs.joinmastodon.org/methods/announcements/#put-reactions
type PutAnnouncement_Reaction struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
	Name string `param:"name"`
}

// DeleteAnnouncement_Reaction is the input for DELETE /api/v1/announcements/:id/reactions/:name, which returns an empty object.
// https://docs.joinmastodon.org/methods/announcements/#delete-reactions
type DeleteAnnouncement_Reaction struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
	Name string `param:"name"`
}
