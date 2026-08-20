package route

/******************************************
 * Follow Requests API Methods
 * View and manage follow requests
 * https://docs.joinmastodon.org/methods/follow_requests/
 ******************************************/

// GetFollowRequests is the route for GET /api/v1/follow_requests.
// https://docs.joinmastodon.org/methods/follow_requests/#get
const GetFollowRequests = "/api/v1/follow_requests"

// PostFollowRequest_Authorize is the route for POST /api/v1/follow_requests/:account_id/authorize.
// https://docs.joinmastodon.org/methods/follow_requests/#accept
const PostFollowRequest_Authorize = "/api/v1/follow_requests/:account_id/authorize"

// PostFollowRequest_Reject is the route for POST /api/v1/follow_requests/:account_id/reject.
// https://docs.joinmastodon.org/methods/follow_requests/#reject
const PostFollowRequest_Reject = "/api/v1/follow_requests/:account_id/reject"
