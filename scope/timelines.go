package scope

/******************************************
 * Timelines API Methods
 * Read and view timelines of statuses
 * https://docs.joinmastodon.org/methods/timelines/
 ******************************************/

// GetTimeline_Public is the OAuth scope required by GET /api/v1/timelines/public.
// https://docs.joinmastodon.org/methods/timelines/#public
const GetTimeline_Public = ReadStatuses

// GetTimeline_Hashtag is the OAuth scope required by GET /api/v1/timelines/tag/:hashtag.
// https://docs.joinmastodon.org/methods/timelines/#tag
const GetTimeline_Hashtag = ReadStatuses

// GetTimeline_Home is the OAuth scope required by GET /api/v1/timelines/home.
// https://docs.joinmastodon.org/methods/timelines/#home
const GetTimeline_Home = ReadStatuses

// GetTimeline_List is the OAuth scope required by GET /api/v1/timelines/list/:list_id.
// https://docs.joinmastodon.org/methods/timelines/#list
const GetTimeline_List = ReadStatuses
