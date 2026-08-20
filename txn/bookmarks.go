package txn

/******************************************
 * Bookmarks API Methods
 * View your bookmarks. See also statuses/:id/(bookmark,unbookmark)
 * https://docs.joinmastodon.org/methods/bookmarks/
 ******************************************/

// GetBookmarks is the input for GET /api/v1/bookmarks, which returns []Status.
// https://docs.joinmastodon.org/methods/bookmarks/#get
type GetBookmarks struct {
	Host    string `header:"Host"`
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	MinID   string `query:"min_id"`
	Limit   int64  `query:"limit"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetBookmarks) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}
