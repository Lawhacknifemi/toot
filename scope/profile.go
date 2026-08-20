package scope

/******************************************
 * Profile API Methods
 * Methods concerning profiles
 * https://docs.joinmastodon.org/methods/profile/
 ******************************************/

// DeleteProfile_Avatar is the OAuth scope required by DELETE /api/v1/profile/avatar.
// https://docs.joinmastodon.org/methods/profile/#delete-profile-avatar
const DeleteProfile_Avatar = WriteAccounts

// DeleteProfile_Header is the OAuth scope required by DELETE /api/v1/profile/header.
// https://docs.joinmastodon.org/methods/profile/#delete-profile-header
const DeleteProfile_Header = WriteAccounts
