package txn

/******************************************
 * Emails API Methods
 * Request a new confirmation email, potentially to a new email address
 * https://docs.joinmastodon.org/methods/emails/
 ******************************************/

// PostEmailConfirmation is the input for POST /api/v1/email/confirmation, which returns an empty object.
// https://docs.joinmastodon.org/methods/emails/#confirmation
type PostEmailConfirmation struct {
	Host string `header:"Host"`
}
