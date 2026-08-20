package scope

/******************************************
 * Domain Blocks API Methods
 * Manage a User's blocked domains.
 * https://docs.joinmastodon.org/methods/domain_blocks/
 ******************************************/

// GetDomainBlocks is the OAuth scope required by GET /api/v1/domain_blocks.
// https://docs.joinmastodon.org/methods/domain_blocks/#get
const GetDomainBlocks = ReadBlocks

// PostDomainBlock is the OAuth scope required by POST /api/v1/domain_blocks.
// https://docs.joinmastodon.org/methods/domain_blocks/#block
const PostDomainBlock = WriteBlocks

// DeleteDomainBlock is the OAuth scope required by DELETE /api/v1/domain_blocks.
// https://docs.joinmastodon.org/methods/domain_blocks/#unblock
const DeleteDomainBlock = WriteBlocks
