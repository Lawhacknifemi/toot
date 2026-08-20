package route

/******************************************
 * Scheduled Statuses API Methods
 * Manage statuses that were scheduled to be published at a future date.
 * https://docs.joinmastodon.org/methods/scheduled_statuses/
 ******************************************/

// GetScheduledStatuses is the route for GET /api/v1/scheduled_statuses.
// https://docs.joinmastodon.org/methods/scheduled_statuses/#get
const GetScheduledStatuses = "/api/v1/scheduled_statuses"

// GetScheduledStatus is the route for GET /api/v1/scheduled_statuses/:id.
// https://docs.joinmastodon.org/methods/scheduled_statuses/#get-one
const GetScheduledStatus = "/api/v1/scheduled_statuses/:id"

// PutScheduledStatus is the route for PUT /api/v1/scheduled_statuses/:id.
// https://docs.joinmastodon.org/methods/scheduled_statuses/#update
const PutScheduledStatus = "/api/v1/scheduled_statuses/:id"

// DeleteScheduledStatus is the route for DELETE /api/v1/scheduled_statuses/:id.
// https://docs.joinmastodon.org/methods/scheduled_statuses/#cancel
const DeleteScheduledStatus = "/api/v1/scheduled_statuses/:id"
