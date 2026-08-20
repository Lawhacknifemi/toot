package route

/******************************************
* Polls API Methods
* View and vode on polls attached to statuses. To discover pollID,
* you will need to GET a Status first and then check for a `poll` property.
* https://docs.joinmastodon.org/methods/polls/
******************************************/

// GetPoll is the route for GET /api/v1/polls/:id.
// https://docs.joinmastodon.org/methods/polls/#get
const GetPoll = "/api/v1/polls/:id"

// PostPoll_Votes is the route for POST /api/v1/polls/:id/votes.
// https://docs.joinmastodon.org/methods/polls/#vote
const PostPoll_Votes = "/api/v1/polls/:id/votes"
