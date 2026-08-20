package txn

/******************************************
 * OEmbed API Methods
 * For generating OEmbed previews
 * https://docs.joinmastodon.org/methods/oembed/
 ******************************************/

// GetOEmbed is the input for GET /api/oembed, which returns OEmbed metadata.
// https://docs.joinmastodon.org/methods/oembed/#get
type GetOEmbed struct {
	Host      string `header:"Host"`
	URL       string `query:"url"`
	MaxWidth  int    `query:"maxwidth"`
	MaxHeight int    `query:"maxheight"`
}
