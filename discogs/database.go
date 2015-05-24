package discogs

// Services housed in the database API https://www.discogs.com/developers/#page:database
// Release, MasterRelease, Master Release Versions. Artist, Artist Releases, Label, All Label Releases

type DatabaseService struct {
	client *Client
}
