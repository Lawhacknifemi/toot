package scope

/******************************************
 * Favourites API Methods
 * View your favourites. See also /statuses/:id/{favourite,unfavourite}
 * https://docs.joinmastodon.org/methods/favourites/
 ******************************************/

// GetFavourites is the OAuth scope required by GET /api/v1/favourites.
// https://docs.joinmastodon.org/methods/favourites/#get
const GetFavourites = ReadFavourites
