package scope

/******************************************
 * Bookmarks API Methods
 * View your bookmarks. See also statuses/:id/(bookmark,unbookmark)
 * https://docs.joinmastodon.org/methods/bookmarks/
 ******************************************/

// GetBookmarks is the OAuth scope required by GET /api/v1/bookmarks.
// https://docs.joinmastodon.org/methods/bookmarks/#get
const GetBookmarks = ReadBookmarks
