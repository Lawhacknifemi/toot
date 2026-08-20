package txn

/******************************************
 * Directory API Methods
 * A directory of profiles that your website is aware of.
 * https://docs.joinmastodon.org/methods/directory/
 ******************************************/

// GetDirectory is the input for GET /api/v1/directory, which returns []Account.
// https://docs.joinmastodon.org/methods/directory/#get
type GetDirectory struct {
	Host   string `header:"Host"`
	Offset int    `query:"offset"`
	Limit  int    `query:"limit"`
	Order  string `query:"order"`
	Local  bool   `query:"local"`
}
