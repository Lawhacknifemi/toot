package scope

/******************************************
 * Scheduled Statuses API Methods
 * Manage statuses that were scheduled to be published at a future date.
 * https://docs.joinmastodon.org/methods/scheduled_statuses/
 ******************************************/

// GetScheduledStatuses is the OAuth scope required by GET /api/v1/scheduled_statuses.
// https://docs.joinmastodon.org/methods/scheduled_statuses/#get
const GetScheduledStatuses = ReadStatuses

// GetScheduledStatus is the OAuth scope required by GET /api/v1/scheduled_statuses/:id.
// https://docs.joinmastodon.org/methods/scheduled_statuses/#get-one
const GetScheduledStatus = ReadStatuses

// PutScheduledStatus is the OAuth scope required by PUT /api/v1/scheduled_statuses/:id.
// https://docs.joinmastodon.org/methods/scheduled_statuses/#update
const PutScheduledStatus = WriteStatuses

// DeleteScheduledStatus is the OAuth scope required by DELETE /api/v1/scheduled_statuses/:id.
// https://docs.joinmastodon.org/methods/scheduled_statuses/#cancel
const DeleteScheduledStatus = WriteStatuses
