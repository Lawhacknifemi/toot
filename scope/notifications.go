package scope

/******************************************
 * Notifications API Methods
 * Receive notifications for activity on your account or statuses
 * https://docs.joinmastodon.org/methods/notifications/
 ******************************************/

// GetNotifications is the OAuth scope required by GET /api/v1/notifications.
// https://docs.joinmastodon.org/methods/notifications/#get
const GetNotifications = ReadNotifications

// GetNotification is the OAuth scope required by GET /api/v1/notifications/:id.
// https://docs.joinmastodon.org/methods/notifications/#get-one
const GetNotification = ReadNotifications

// PostNotifications_Clear is the OAuth scope required by POST /api/v1/notifications/clear.
// https://docs.joinmastodon.org/methods/notifications/#clear
const PostNotifications_Clear = WriteNotifications

// PostNotification_Dismiss is the OAuth scope required by POST /api/v1/notifications/dismiss.
// https://docs.joinmastodon.org/methods/notifications/#dismiss
const PostNotification_Dismiss = WriteNotifications
