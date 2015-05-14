package discogs

// Common pagination response, used within AlbumReleases and LabelReleases
type Paginate struct {
	PerPage int `json:"per_page"`
	Pages   int `json:"pages"`
	Page    int `json:"page"`
	Urls    struct {
		Last string `json:"last"`
		Next string `json:"next"`
	} `json:"urls"`
	Items int `json:"items"`
}
