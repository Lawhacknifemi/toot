package object

// StatusMention represents an account mentioned within the content of a status.
// https://docs.joinmastodon.org/entities/Status/#mentions
type StatusMention struct {
	ID       string `json:"id"`       // The account ID of the mentioned user.
	Username string `json:"username"` // The username of the mentioned user.
	URL      string `json:"url"`      // The location of the mentioned user’s profile.
	Acct     string `json:"acct"`     // The webfinger acct: URI of the mentioned user. Equivalent to username for local users, or username@domain for remote users.
}
