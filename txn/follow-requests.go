package txn

/******************************************
 * Follow Requests API Methods
 * View and manage follow requests
 * https://docs.joinmastodon.org/methods/follow_requests/
 ******************************************/

// GetFollowRequests is the input for GET /api/v1/follow_requests, which returns []Account.
// https://docs.joinmastodon.org/methods/follow_requests/#get
type GetFollowRequests struct {
	Host    string `header:"Host"`
	MaxID   string `query:"max_id"`
	MinID   string `query:"min_id"`
	SinceID string `query:"since_id"`
	Limit   int64  `query:"limit"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetFollowRequests) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}

// PostFollowRequest_Authorize is the input for POST /api/v1/follow_requests/:account_id/authorize, which returns Relationship.
// https://docs.joinmastodon.org/methods/follow_requests/#accept
type PostFollowRequest_Authorize struct {
	Host      string `header:"Host"`
	AccountID string `param:"account_id"`
}

// PostFollowRequest_Reject is the input for POST /api/v1/follow_requests/:account_id/reject, which returns Relationship.
// https://docs.joinmastodon.org/methods/follow_requests/#reject
type PostFollowRequest_Reject struct {
	Host      string `header:"Host"`
	AccountID string `param:"account_id"`
}
