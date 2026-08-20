package scope

/******************************************
 * Conversations API Methods
 * Direct conversations with other participants.
 * (Currently, just threads containing a post with "direct" visibility.)
 * https://docs.joinmastodon.org/methods/conversations/
 ******************************************/

// GetConversations is the OAuth scope required by GET /api/v1/conversations.
// https://docs.joinmastodon.org/methods/conversations/#get
const GetConversations = ReadStatuses

// DeleteConversation is the OAuth scope required by DELETE /api/v1/conversations/:id.
// https://docs.joinmastodon.org/methods/conversations/#delete
const DeleteConversation = WriteConversations

// PostConversationRead is the OAuth scope required by POST /api/v1/conversations/:id/read.
// https://docs.joinmastodon.org/methods/conversations/#read
const PostConversationRead = WriteConversations
