package object

// List represents a list of some users that the authenticated user follows.
// https://docs.joinmastodon.org/entities/List/
type List struct {
	ID            string `json:"id"`             // The internal database ID of the list.
	Title         string `json:"title"`          // The user-defined title of the list.
	RepliesPolicy string `json:"replies_policy"` // Which replies should be shown in the list. [followed | list | none]
}

// ListRepliesPolicyFollowed shows replies to any followed user.
const ListRepliesPolicyFollowed = "followed"

// ListRepliesPolicyList shows replies to members of the list.
const ListRepliesPolicyList = "list"

// ListRepliesPolicyNone shows replies to no one.
const ListRepliesPolicyNone = "none"
