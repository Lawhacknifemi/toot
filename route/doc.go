// Package route contains the URL path of every Mastodon API endpoint, as an
// Echo-style pattern with ":name" placeholders for the path parameters.
//
// Each constant is named for the endpoint it serves -- GetNotifications,
// PostAccount_Follow -- and the same name identifies that endpoint in the toot,
// scope, and txn packages, so a router can pair a path with its required scope
// and its input type without a lookup table.
//
// Paths for the admin API, push subscriptions, and streaming are not defined yet.
package route
