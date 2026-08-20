package object

// TagHistory holds the usage statistics for a hashtag on a given day.
// https://docs.joinmastodon.org/entities/Tag/#history
type TagHistory struct {
	Day      string `json:"day"`      // UNIX timestamp on midnight of the given day.
	Uses     string `json:"uses"`     // The counted usage of the tag within that day.
	Accounts string `json:"accounts"` // The total of accounts using the tag within that day.
}
