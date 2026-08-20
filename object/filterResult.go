package object

// FilterResult represents a filter that matched a status, and the keywords that matched it.
// https://docs.joinmastodon.org/entities/FilterResult/
type FilterResult struct {
	Filter         Filter   `json:"filter"`                    // The filter that was matched.
	KeywordMatches []string `json:"keyword_matches,omitempty"` // The keyword within the filter that was matched.
	StatusMatches  []string `json:"status_matches,omitempty"`  // The status ID within the filter that was matched.
}
