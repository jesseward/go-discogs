package discogs

import (
	"fmt"
)

type Label struct {
	ID          int        `json:"id"`
	Profile     string     `json:"profile"`
	ReleasesURL string     `json:"releases_url"`
	Name        string     `json:"name"`
	ContactInfo string     `json:"contact_info"`
	URI         string     `json:"uri"`
	Sublabels   []Sublabel `json:"sublabels"`
	Urls        []string   `json:"urls"`
	Images      []Image    `json:"images"`
	ResourceURL string     `json:"resource_url"`
	DataQuality string     `json:"data_quality"`
}

type Sublabel struct {
	ResourceURL string `json:"resource_url"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
}

type LabelService struct {
	client *Client
}

type LabelReleases struct {
	Pagination Paginate `json:"pagination"`
	Releases   []struct {
		Artist      string `json:"artist"`
		Catno       string `json:"catno"`
		Format      string `json:"format"`
		ID          int    `json:"id"`
		ResourceURL string `json:"resource_url"`
		Status      string `json:"status"`
		Thumb       string `json:"thumb"`
		Title       string `json:"title"`
		Year        int    `json:"year"`
	} `json:"releases"`
}

func (s *LabelService) Releases(id int) (*LabelReleases, *Response, error) {
	url := fmt.Sprintf("labels/%d/releases", id)

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	label := new(LabelReleases)
	resp, err := s.client.Do(req, label)

	if err != nil {
		return nil, resp, err
	}

	return label, resp, err
}

func (s *LabelService) Get(id int) (*Label, *Response, error) {
	url := fmt.Sprintf("labels/%d", id)

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	label := new(Label)
	resp, err := s.client.Do(req, label)

	if err != nil {
		return nil, resp, err
	}

	return label, resp, err
}
