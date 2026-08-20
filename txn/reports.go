package txn

/******************************************
 * Reports API Methods
 * Report problematic users to your moderators
 * https://docs.joinmastodon.org/methods/reports/
 ******************************************/

// PostReport is the input for POST /api/v1/reports, which returns Report.
// https://docs.joinmastodon.org/methods/reports/#post
type PostReport struct {
	Host      string   `header:"Host"`
	AccountID string   `form:"account_id"`
	StatusIDs []string `form:"status_ids"`
	Comment   string   `form:"comment"`
	Forward   bool     `form:"forward"`
	Category  string   `form:"category"`
	RuleIDs   []string `form:"rule_ids"`
}
