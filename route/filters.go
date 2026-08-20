package route

/******************************************
 * Filters API Methods
 * Create and manage filters
 * https://docs.joinmastodon.org/methods/filters/
 ******************************************/

// GetFilters is the route for GET /api/v2/filters.
// https://docs.joinmastodon.org/methods/filters/#get
const GetFilters = "/api/v2/filters"

// GetFilter is the route for GET /api/v2/filters/:id.
// https://docs.joinmastodon.org/methods/filters/#get-one
const GetFilter = "/api/v2/filters/:id"

// PostFilter is the route for POST /api/v2/filters.
// https://docs.joinmastodon.org/methods/filters/#create
const PostFilter = "/api/v2/filters"

// PutFilter is the route for PUT /api/v2/filters/:id.
// https://docs.joinmastodon.org/methods/filters/#update
const PutFilter = "/api/v2/filters/:id"

// DeleteFilter is the route for DELETE /api/v2/filters/:id.
// https://docs.joinmastodon.org/methods/filters/#delete
const DeleteFilter = "/api/v2/filters/:id"

// GetFilter_Keywords is the route for GET /api/v2/filters/:filter_id/keywords.
// https://docs.joinmastodon.org/methods/filters/#keywords-get
const GetFilter_Keywords = "/api/v2/filters/:filter_id/keywords"

// PostFilter_Keyword is the route for POST /api/v2/filters/:filter_id/keywords.
// https://docs.joinmastodon.org/methods/filters/#keywords-create
const PostFilter_Keyword = "/api/v2/filters/:filter_id/keywords"

// GetFilter_Keyword is the route for GET /api/v2/filters/keywords/:id.
// https://docs.joinmastodon.org/methods/filters/#keywords-get-one
const GetFilter_Keyword = "/api/v2/filters/keywords/:id"

// PutFilter_Keyword is the route for PUT /api/v2/filters/keywords/:id.
// https://docs.joinmastodon.org/methods/filters/#keywords-update
const PutFilter_Keyword = "/api/v2/filters/keywords/:id"

// DeleteFilter_Keyword is the route for DELETE /api/v2/filters/keywords/:id.
// https://docs.joinmastodon.org/methods/filters/#keywords-delete
const DeleteFilter_Keyword = "/api/v2/filters/keywords/:id"

// GetFilter_Statuses is the route for GET /api/v2/filters/:filter_id/statuses.
// https://docs.joinmastodon.org/methods/filters/#statuses-get
const GetFilter_Statuses = "/api/v2/filters/:filter_id/statuses"

// PostFilter_Status is the route for POST /api/v2/filters/:filter_id/statuses.
// https://docs.joinmastodon.org/methods/filters/#statuses-add
const PostFilter_Status = "/api/v2/filters/:filter_id/statuses"

// GetFilter_Status is the route for GET /api/v2/filters/statuses/:id.
// https://docs.joinmastodon.org/methods/filters/#statuses-get-one
const GetFilter_Status = "/api/v2/filters/statuses/:id"

// DeleteFilter_Status is the route for DELETE /api/v2/filters/statuses/:id.
// https://docs.joinmastodon.org/methods/filters/#statuses-remove
const DeleteFilter_Status = "/api/v2/filters/statuses/:id"

// GetFilters_V1 is the route for GET /api/v1/filters.
// https://docs.joinmastodon.org/methods/filters/#get-v1
const GetFilters_V1 = "/api/v1/filters"

// GetFilter_V1 is the route for GET /api/v1/filters/:id.
// https://docs.joinmastodon.org/methods/filters/#get-one-v1
const GetFilter_V1 = "/api/v1/filters/:id"

// PostFilter_V1 is the route for POST /api/v1/filters.
// https://docs.joinmastodon.org/methods/filters/#create-v1
const PostFilter_V1 = "/api/v1/filters"

// PutFilter_V1 is the route for PUT /api/v1/filters/:id.
// https://docs.joinmastodon.org/methods/filters/#update-v1
const PutFilter_V1 = "/api/v1/filters/:id"

// DeleteFilter_V1 is the route for DELETE /api/v1/filters/:id.
// https://docs.joinmastodon.org/methods/filters/#delete-v1
const DeleteFilter_V1 = "/api/v1/filters/:id"
