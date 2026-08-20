package scope

/******************************************
 * Markers API Methods
 * Save and restore your position in timelines
 * https://docs.joinmastodon.org/methods/markers/
 ******************************************/

// GetMarkers is the OAuth scope required by GET /api/v1/markers.
// https://docs.joinmastodon.org/methods/markers/#get
const GetMarkers = ReadStatuses

// PostMarker is the OAuth scope required by POST /api/v1/markers.
// https://docs.joinmastodon.org/methods/markers/#create
const PostMarker = WriteStatuses
