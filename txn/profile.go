package txn

/******************************************
 * Profile API Methods
 * Methods concerning profiles
 * https://docs.joinmastodon.org/methods/profile/
 ******************************************/

// DeleteProfile_Avatar is the input for DELETE /api/v1/profile/avatar, which returns CredentialAccount.
// https://docs.joinmastodon.org/methods/profile/#delete-profile-avatar
type DeleteProfile_Avatar struct {
	Host string `header:"Host"`
}

// DeleteProfile_Header is the input for DELETE /api/v1/profile/header, which returns CredentialAccount.
// https://docs.joinmastodon.org/methods/profile/#delete-profile-header
type DeleteProfile_Header struct {
	Host string `header:"Host"`
}
