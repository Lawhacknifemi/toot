// Package txn contains the input type of every Mastodon API endpoint.
//
// Each type collects the arguments of one API call, with struct tags naming
// where each field is read from: "param" for a path placeholder, "query" for a
// query-string value, "form" for a posted form field, and "header" for a request
// header. A binder fills the struct from an http.Request before the handler runs.
//
// Each type is named for the endpoint it serves -- GetNotifications,
// PostAccount_Follow -- matching the constants in the route and scope packages
// and the handler fields on toot.API.
//
// Types that accept pagination embed the same four parameters, and report them
// through the QueryPager interface.
//
// Input types for the admin API, push subscriptions, and streaming are not
// defined yet.
package txn
