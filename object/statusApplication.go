package object

// StatusApplication represents the application used to post a status.
// https://docs.joinmastodon.org/entities/Status/#application
type StatusApplication struct {
	Name    string `json:"name"`              // The name of the application that posted this status.
	Website string `json:"website,omitempty"` // The website associated with the application that posted this status.
}
