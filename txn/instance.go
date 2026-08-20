package txn

/******************************************
 * Instance API Methods
 * Discover information about a Mastodon website
 * https://docs.joinmastodon.org/methods/instance/
 ******************************************/

// GetInstance is the input for GET /api/v2/instance, which returns Instance.
// https://docs.joinmastodon.org/methods/instance/#v2
type GetInstance struct {
	Host string `header:"Host"`
}

// GetInstance_Peers is the input for GET /api/v1/instance/peers, which returns []string.
// https://docs.joinmastodon.org/methods/instance/#peers
type GetInstance_Peers struct {
	Host string `header:"Host"`
}

// GetInstance_Activity is the input for GET /api/v1/instance/activity, which returns []map[string]any.
// https://docs.joinmastodon.org/methods/instance/#activity
type GetInstance_Activity struct {
	Host string `header:"Host"`
}

// GetInstance_Rules is the input for GET /api/v1/instance/rules, which returns []Rule.
// https://docs.joinmastodon.org/methods/instance/#rules
type GetInstance_Rules struct {
	Host string `header:"Host"`
}

// GetInstance_DomainBlocks is the input for GET /api/v1/instance/domain_blocks, which returns []DomainBlock.
// https://docs.joinmastodon.org/methods/instance/#domain_blocks
type GetInstance_DomainBlocks struct {
	Host string `header:"Host"`
}

// GetInstance_ExtendedDescription is the input for GET /api/v1/instance/extended_description, which returns ExtendedDescription.
// https://docs.joinmastodon.org/methods/instance/#extended_description
type GetInstance_ExtendedDescription struct {
	Host string `header:"Host"`
}

// GetInstance_V1 is the input for GET /api/v1/instance, which returns Instance_V1.
// https://docs.joinmastodon.org/methods/instance/#v1
type GetInstance_V1 struct {
	Host string `header:"Host"`
}
