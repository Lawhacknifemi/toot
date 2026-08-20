package txn

/******************************************
 * Favourites API Methods
 * View your favourites. See also /statuses/:id/{favourite,unfavourite}
 * https://docs.joinmastodon.org/methods/favourites/
 ******************************************/

// GetFavourites is the input for GET /api/v1/favourites, which returns []Status.
// https://docs.joinmastodon.org/methods/favourites/#get
type GetFavourites struct {
	Host    string `header:"Host"`
	MaxID   string `query:"max_id"`
	SinceID string `query:"since_id"`
	MinID   string `query:"min_id"`
	Limit   int64  `query:"limit"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetFavourites) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}
