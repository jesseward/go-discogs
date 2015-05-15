package discogs

import (
	"fmt"
)

type ArtistService struct {
	client *Client
}

// https://www.discogs.com/developers/#page:database,header:database-artist
type Artist struct {
	ID             int      `json:"id"`
	Profile        string   `json:"profile"`
	Realname       string   `json:"realname"`
	ReleasesURL    string   `json:"releases_url"`
	Members        []Member `json:"members"`
	Name           string   `json:"name"`
	URI            string   `json:"uri"`
	URLS           []string `json:"urls""`
	Images         []Image  `json:"images"`
	ResourceURL    string   `json:"resource_url"`
	Aliases        []Alias  `json:"aliases"`
	DataQuality    string   `json:"data_quality"`
	Namevariations []string `json:"namevariations"`
}

type Alias struct {
	ResourceURL string `json:"resource_url"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
}

type Member struct {
	Active      bool   `json:"active"`
	ResourceURL string `json:"resource_url"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
}

// https://www.discogs.com/developers/#page:database,header:database-artist-releases
type ArtistReleases struct {
	Pagination Paginate `json:"pagination"`
	Releases   []struct {
		Artist      string `json:"artist"`
		ID          int    `json:"id"`
		MainRelease int    `json:"main_release"`
		ResourceURL string `json:"resource_url"`
		Role        string `json:"role"`
		Thumb       string `json:"thumb"`
		Title       string `json:"title"`
		Type        string `json:"type"`
		Year        int    `json:"year"`
	} `json:"releases"`
}

func (s *ArtistService) GetReleases(id int) (*ArtistReleases, *Response, error) {

	url := fmt.Sprintf("artists/%d/releases", id)

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	artist := new(ArtistReleases)
	resp, err := s.client.Do(req, artist)

	if err != nil {
		return nil, resp, err
	}

	return artist, resp, err
}

func (s *ArtistService) Get(id int) (*Artist, *Response, error) {
	url := fmt.Sprintf("artists/%d", id)

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	artist := new(Artist)
	resp, err := s.client.Do(req, artist)

	if err != nil {
		return nil, resp, err
	}

	return artist, resp, err
}
