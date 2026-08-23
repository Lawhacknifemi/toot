package txn

/******************************************
 * Notifications API Methods
 * Receive notifications for activity on your account or statuses
 * https://docs.joinmastodon.org/methods/notifications/
 ******************************************/

// GetNotifications is the input for GET /api/v1/notifications, which returns []Notification.
// https://docs.joinmastodon.org/methods/notifications/#get
type GetNotifications struct {
	Host         string   `header:"Host"`
	MaxID        string   `query:"max_id"`
	SinceID      string   `query:"since_id"`
	MinID        string   `query:"min_id"`
	Limit        int64    `query:"limit"`
	Types        []string `query:"types"`
	ExcludeTypes []string `query:"exclude_types"`
	AccountID    string   `query:"account_id"`
}

// QueryPage implements the QueryPager interface, returning
// the QueryPage data embedded in this transaction
func (t GetNotifications) QueryPage() QueryPage {
	return QueryPage{
		MaxID:   t.MaxID,
		SinceID: t.SinceID,
		MinID:   t.MinID,
		Limit:   t.Limit,
	}
}

// GetNotification is the input for GET /api/v1/notifications/:id, which returns Notification.
// https://docs.joinmastodon.org/methods/notifications/#get-one
type GetNotification struct {
	Host string `header:"Host"`
	ID   string `param:"id"`
}

// PostNotifications_Clear is the input for POST /api/v1/notifications/clear, which returns an empty object.
// https://docs.joinmastodon.org/methods/notifications/#clear
type PostNotifications_Clear struct {
	Host string `header:"Host"`
}

// PostNotification_Dismiss is the input for POST /api/v1/notifications/dismiss, which returns an empty object.
// https://docs.joinmastodon.org/methods/notifications/#dismiss
type PostNotification_Dismiss struct {
	Host string `header:"Host"`
	ID   string `form:"id"`
}

// GetNotifications_UnreadCount is the input for GET /api/v1/notifications/unread_count, which returns NotificationsUnreadCount.
// https://docs.joinmastodon.org/methods/notifications/#unread-count
type GetNotifications_UnreadCount struct {
	Host         string   `header:"Host"`
	Limit        int      `query:"limit"`
	Types        []string `query:"types"`
	ExcludeTypes []string `query:"exclude_types"`
	AccountID    string   `query:"account_id"`
}

// GetNotificationPolicy is the input for GET /api/v2/notifications/policy, which returns NotificationPolicy.
// https://docs.joinmastodon.org/methods/notifications/#get-policy
type GetNotificationPolicy struct {
	Host string `header:"Host"`
}

// PatchNotificationPolicy is the input for PATCH /api/v2/notifications/policy, which returns NotificationPolicy.
// https://docs.joinmastodon.org/methods/notifications/#update-policy
type PatchNotificationPolicy struct {
	Host               string `header:"Host"`
	ForNotFollowing    string `form:"for_not_following"`
	ForNotFollowers    string `form:"for_not_followers"`
	ForNewAccounts     string `form:"for_new_accounts"`
	ForPrivateMentions string `form:"for_private_mentions"`
	ForLimitedAccounts string `form:"for_limited_accounts"`
	ForBots            string `form:"for_bots"`
}
