package discogs

import (
	"fmt"
)

// https://www.discogs.com/developers/#page:database,header:database-release
type Release struct {
	Title             string           `json:"title"`
	ID                int              `json:"id"`
	Artists           []ArtistResource `json:"artists"`
	DataQuality       string           `json:"data_quality"`
	Thumb             string           `json:"thumb"`
	Community         Community        `json:"community"`
	Companies         []Company        `json:"companies"`
	Country           string           `json:"country"`
	DateAdded         string           `json:"date_added"`
	DateChanged       string           `json:"date_changed"`
	EstimatedWeight   int              `json:"estimated_weight"`
	Extraartists      []ArtistResource `json:"extraartists"`
	FormatQuantity    int              `json:"format_quantity"`
	Formats           []Format         `json:"formats"`
	Genres            []string         `json:"genres"`
	Identifiers       []Identifier     `json:"identifiers"`
	Images            []Image          `json:"images"`
	Labels            []ReleaseLabel   `json:"labels"`
	MasterID          int              `json:"master_id"`
	MasterURL         string           `json:"master_url"`
	Notes             string           `json:"notes"`
	Released          string           `json:"released"`
	ReleasedFormatted string           `json:"released_formatted"`
	ResourceURL       string           `json:"resource_url"`
	Series            []interface{}    `json:"series"`
	Status            string           `json:"status"`
	Styles            []string         `json:"styles"`
	Tracklist         []Track          `json:"tracklist"`
	URI               string           `json:"uri"`
	Videos            []Video          `json:"videos"`
	Year              int              `json:"year"`
}

type ArtistResource struct {
	Join        string `json:"join"`
	Name        string `json:"name"`
	Anv         string `json:"anv"`
	Tracks      string `json:"tracks"`
	Role        string `json:"role"`
	ResourceURL string `json:"resource_url"`
	ID          int    `json:"id"`
}

type Company struct {
	Name           string `json:"name"`
	EntityType     string `json:"entity_type"`
	Catno          string `json:"catno"`
	ResourceURL    string `json:"resource_url"`
	ID             int    `json:"id"`
	EntityTypeName string `json:"entity_type_name"`
}

type Format struct {
	Qty          string   `json:"qty"`
	Descriptions []string `json:"descriptions"`
	Name         string   `json:"name"`
}

type Identifier struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Value       string `json:"value"`
}

type ReleaseLabel struct {
	ResourceURL string `json:"resource_url"`
	EntityType  string `json:"entity_type"`
	Catno       string `json:"catno"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
}

type Track struct {
	Duration string `json:"duration"`
	Position string `json:"position"`
	Type     string `json:"type_"`
	Title    string `json:"title"`
}
type Video struct {
	Duration    int    `json:"duration"`
	Embed       bool   `json:"embed"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URI         string `json:"uri"`
}

type ReleaseService struct {
	client *Client
}

func (s *ReleaseService) Get(id int) (*Release, *Response, error) {
	url := fmt.Sprintf("releases/%d", id)

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	release := new(Release)
	resp, err := s.client.Do(req, release)

	if err != nil {
		return nil, resp, err
	}

	return release, resp, err
}
