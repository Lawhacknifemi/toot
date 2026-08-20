package txn

/******************************************
* Polls API Methods
* View and vode on polls attached to statuses. To discover pollID,
* you will need to GET a Status first and then check for a `poll` property.
* https://docs.joinmastodon.org/methods/polls/
******************************************/

// GetPoll is the input for GET /api/v1/polls/:id, which returns Poll.
// https://docs.joinmastodon.org/methods/polls/#get
type GetPoll struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostPoll_Votes is the input for POST /api/v1/polls/:id/votes, which returns Poll.
// https://docs.joinmastodon.org/methods/polls/#vote
type PostPoll_Votes struct {
	Host    string `header:"Host"`
	ID      string `param:"id"`
	Choices []int  `json:"choices"`
}
