package scope

/******************************************
 * Endorsements API Methods
 * Feature other profiles on your own profile. See also accounts/:id/{pin,unpin}
 * https://docs.joinmastodon.org/methods/endorsements/
 ******************************************/

// GetEndorsements is the OAuth scope required by GET /api/v1/endorsements.
// https://docs.joinmastodon.org/methods/endorsements/#get
const GetEndorsements = ReadAccounts
