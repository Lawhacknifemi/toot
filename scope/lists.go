package scope

/******************************************
 * Lists API Methods
 * View and manage lists. See also /api/v1/timelines/list/:id for loading a list timeline
 * https://docs.joinmastodon.org/methods/lists/
 ******************************************/

// GetLists is the OAuth scope required by GET /api/v1/lists.
// https://docs.joinmastodon.org/methods/lists/#get
const GetLists = ReadLists

// GetList is the OAuth scope required by GET /api/v1/lists/:id.
// https://docs.joinmastodon.org/methods/lists/#get-one
const GetList = WriteLists

// PostList is the OAuth scope required by POST /api/v1/lists.
// https://docs.joinmastodon.org/methods/lists/#create
const PostList = WriteLists

// PutList is the OAuth scope required by PUT /api/v1/lists/:id.
// https://docs.joinmastodon.org/methods/lists/#update
const PutList = WriteLists

// DeleteList is the OAuth scope required by DELETE /api/v1/lists/:id.
// https://docs.joinmastodon.org/methods/lists/#delete
const DeleteList = WriteLists

// GetList_Accounts is the OAuth scope required by GET /api/v1/lists/:id/accounts.
// https://docs.joinmastodon.org/methods/lists/#accounts
const GetList_Accounts = ReadLists

// PostList_Accounts is the OAuth scope required by POST /api/v1/lists/:id/accounts.
// https://docs.joinmastodon.org/methods/lists/#accounts-add
const PostList_Accounts = WriteLists

// DeleteList_Accounts is the OAuth scope required by DELETE /api/v1/lists/:id/accounts.
// https://docs.joinmastodon.org/methods/lists/#accounts-remove
const DeleteList_Accounts = WriteLists
