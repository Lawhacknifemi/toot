package scope

/******************************************
 * Featured Tags API Methods
 * Feature tags that you use frequently on your profile
 * https://docs.joinmastodon.org/methods/featured_tags/
 ******************************************/

// GetFeaturedTags is the OAuth scope required by GET /api/v1/featured_tags.
// https://docs.joinmastodon.org/methods/featured_tags/#get
const GetFeaturedTags = ReadAccounts

// PostFeaturedTag is the OAuth scope required by POST /api/v1/featured_tags.
// https://docs.joinmastodon.org/methods/featured_tags/#feature
const PostFeaturedTag = WriteAccounts

// DeleteFeaturedTag is the OAuth scope required by DELETE /api/v1/featured_tags/:id.
// https://docs.joinmastodon.org/methods/featured_tags/#unfeature
const DeleteFeaturedTag = WriteAccounts

// GetFeaturedTags_Suggestions is the OAuth scope required by GET /api/v1/featured_tags/suggestions.
// https://docs.joinmastodon.org/methods/featured_tags/#suggestions
const GetFeaturedTags_Suggestions = ReadAccounts
