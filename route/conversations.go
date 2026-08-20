package route

/******************************************
 * Conversations API Methods
 * Direct conversations with other participants.
 * (Currently, just threads containing a post with "direct" visibility.)
 * https://docs.joinmastodon.org/methods/conversations/
 ******************************************/

// GetConversations is the route for GET /api/v1/conversations.
// https://docs.joinmastodon.org/methods/conversations/#get
const GetConversations = "/api/v1/conversations"

// DeleteConversation is the route for DELETE /api/v1/conversations/:id.
// https://docs.joinmastodon.org/methods/conversations/#delete
const DeleteConversation = "/api/v1/conversations/:id"

// PostConversationRead is the route for POST /api/v1/conversations/:id/read.
// https://docs.joinmastodon.org/methods/conversations/#read
const PostConversationRead = "/api/v1/conversations/:id/read"
