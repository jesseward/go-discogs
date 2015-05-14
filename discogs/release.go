package discogs

import (
	"fmt"
)

// https://www.discogs.com/developers/#page:database,header:database-release
type Release struct {
	Styles            []string         `json:"styles"`
	Videos            []Video          `json:"videos"`
	Series            []interface{}    `json:"series"`
	Labels            []ReleaseLabel   `json:"labels"`
	Community         Community        `json:"community"`
	Year              int              `json:"year"`
	Images            []Image          `json:"images"`
	FormatQuantity    int              `json:"format_quantity"`
	ID                int              `json:"id"`
	Genres            []string         `json:"genres"`
	Thumb             string           `json:"thumb"`
	Extraartists      []ArtistResource `json:"extraartists"`
	Title             string           `json:"title"`
	Artists           []ArtistResource `json:"artists"`
	DateChanged       string           `json:"date_changed"`
	MasterID          int              `json:"master_id"`
	Tracklist         []Track          `json:"tracklist"`
	Status            string           `json:"status"`
	ReleasedFormatted string           `json:"released_formatted"`
	EstimatedWeight   int              `json:"estimated_weight"`
	MasterURL         string           `json:"master_url"`
	Released          string           `json:"released"`
	DateAdded         string           `json:"date_added"`
	Country           string           `json:"country"`
	Notes             string           `json:"notes"`
	Identifiers       []Identifier     `json:"identifiers"`
	Companies         []Company        `json:"companies"`
	URI               string           `json:"uri"`
	Formats           []Format         `json:"formats"`
	ResourceURL       string           `json:"resource_url"`
	DataQuality       string           `json:"data_quality"`
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
