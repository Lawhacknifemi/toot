package scope

/******************************************
 * Filters API Methods
 * Create and manage filters
 * https://docs.joinmastodon.org/methods/filters/
 ******************************************/

// GetFilters is the OAuth scope required by GET /api/v2/filters.
// https://docs.joinmastodon.org/methods/filters/#get
const GetFilters = ReadFilters

// GetFilter is the OAuth scope required by GET /api/v2/filters/:id.
// https://docs.joinmastodon.org/methods/filters/#get-one
const GetFilter = ReadFilters

// PostFilter is the OAuth scope required by POST /api/v2/filters.
// https://docs.joinmastodon.org/methods/filters/#create
const PostFilter = WriteFilters

// PutFilter is the OAuth scope required by PUT /api/v2/filters/:id.
// https://docs.joinmastodon.org/methods/filters/#update
const PutFilter = WriteFilters

// DeleteFilter is the OAuth scope required by DELETE /api/v2/filters/:id.
// https://docs.joinmastodon.org/methods/filters/#delete
const DeleteFilter = WriteFilters

// GetFilter_Keywords is the OAuth scope required by GET /api/v2/filters/:filter_id/keywords.
// https://docs.joinmastodon.org/methods/filters/#keywords-get
const GetFilter_Keywords = ReadFilters

// PostFilter_Keyword is the OAuth scope required by POST /api/v2/filters/:filter_id/keywords.
// https://docs.joinmastodon.org/methods/filters/#keywords-create
const PostFilter_Keyword = WriteFilters

// GetFilter_Keyword is the OAuth scope required by GET /api/v2/filters/keywords/:id.
// https://docs.joinmastodon.org/methods/filters/#keywords-get-one
const GetFilter_Keyword = ReadFilters

// PutFilter_Keyword is the OAuth scope required by PUT /api/v2/filters/keywords/:id.
// https://docs.joinmastodon.org/methods/filters/#keywords-update
const PutFilter_Keyword = WriteFilters

// DeleteFilter_Keyword is the OAuth scope required by DELETE /api/v2/filters/keywords/:id.
// https://docs.joinmastodon.org/methods/filters/#keywords-delete
const DeleteFilter_Keyword = WriteFilters

// GetFilter_Statuses is the OAuth scope required by GET /api/v2/filters/:filter_id/statuses.
// https://docs.joinmastodon.org/methods/filters/#statuses-get
const GetFilter_Statuses = ReadFilters

// PostFilter_Status is the OAuth scope required by POST /api/v2/filters/:filter_id/statuses.
// https://docs.joinmastodon.org/methods/filters/#statuses-add
const PostFilter_Status = WriteFilters

// GetFilter_Status is the OAuth scope required by GET /api/v2/filters/statuses/:id.
// https://docs.joinmastodon.org/methods/filters/#statuses-get-one
const GetFilter_Status = ReadFilters

// DeleteFilter_Status is the OAuth scope required by DELETE /api/v2/filters/statuses/:id.
// https://docs.joinmastodon.org/methods/filters/#statuses-remove
const DeleteFilter_Status = WriteFilters

// GetFilters_V1 is the OAuth scope required by GET /api/v1/filters.
// https://docs.joinmastodon.org/methods/filters/#get-v1
const GetFilters_V1 = ReadFilters

// GetFilter_V1 is the OAuth scope required by GET /api/v1/filters/:id.
// https://docs.joinmastodon.org/methods/filters/#get-one-v1
const GetFilter_V1 = ReadFilters

// PostFilter_V1 is the OAuth scope required by POST /api/v1/filters.
// https://docs.joinmastodon.org/methods/filters/#create-v1
const PostFilter_V1 = WriteFilters

// PutFilter_V1 is the OAuth scope required by PUT /api/v1/filters/:id.
// https://docs.joinmastodon.org/methods/filters/#update-v1
const PutFilter_V1 = WriteFilters

// DeleteFilter_V1 is the OAuth scope required by DELETE /api/v1/filters/:id.
// https://docs.joinmastodon.org/methods/filters/#delete-v1
const DeleteFilter_V1 = WriteFilters
