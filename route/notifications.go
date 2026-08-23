package route

/******************************************
 * Notifications API Methods
 * Receive notifications for activity on your account or statuses
 * https://docs.joinmastodon.org/methods/notifications/
 ******************************************/

// GetNotifications is the route for GET /api/v1/notifications.
// https://docs.joinmastodon.org/methods/notifications/#get
const GetNotifications = "/api/v1/notifications"

// GetNotification is the route for GET /api/v1/notifications/:id.
// https://docs.joinmastodon.org/methods/notifications/#get-one
const GetNotification = "/api/v1/notifications/:id"

// PostNotifications_Clear is the route for POST /api/v1/notifications/clear.
// https://docs.joinmastodon.org/methods/notifications/#clear
const PostNotifications_Clear = "/api/v1/notifications/clear"

// PostNotification_Dismiss is the route for POST /api/v1/notifications/dismiss.
// https://docs.joinmastodon.org/methods/notifications/#dismiss
const PostNotification_Dismiss = "/api/v1/notifications/dismiss"

// https://docs.joinmastodon.org/methods/notifications/#unread-count
const GetNotifications_UnreadCount = "/api/v1/notifications/unread_count"
