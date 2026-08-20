// Package object contains the entities returned by the Mastodon API.
//
// Each type mirrors one entity from the Mastodon documentation -- Account,
// Status, Notification, and the rest -- with JSON tags that match the wire
// format exactly, so a value marshals straight into an API response.
//
// These are data carriers only. They hold no behavior, and validate nothing.
package object
