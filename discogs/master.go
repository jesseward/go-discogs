package discogs

import (
	"fmt"
)

type MasterService struct {
	client *Client
}

// https://www.discogs.com/developers/#page:database,header:database-master-release-versions
type Master struct {
	Artists        []ArtistResource `json:"artists"`
	DataQuality    string           `json:"data_quality"`
	Genres         []string         `json:"genres"`
	ID             int              `json:"id"`
	Images         []Image          `json:"images"`
	MainRelease    int              `json:"main_release"`
	MainReleaseURL string           `json:"main_release_url"`
	Styles         []string         `json:"styles"`
	Videos         []Video          `json:"videos"`
	Title          string           `json:"title"`
	Tracklist      []Track          `json:"tracklist"`
	URI            string           `json:"uri"`
	VersionsURL    string           `json:"versions_url"`
	Year           int              `json:"year"`
	ResourceURL    string           `json:"resource_url"`
}

// https://www.discogs.com/developers/#page:database,header:database-master-release-versions
type MasterVersions struct {
	Pagination Paginate  `json:"pagination"`
	Versions   []Version `json:"versions"`
}

type Version struct {
	Catno       string `json:"catno"`
	Country     string `json:"country"`
	Format      string `json:"format"`
	ID          int    `json:"id"`
	Label       string `json:"label"`
	Released    string `json:"released"`
	ResourceURL string `json:"resource_url"`
	Status      string `json:"status"`
	Thumb       string `json:"thumb"`
	Title       string `json:"title"`
}

func (s *MasterService) Get(id int) (*Master, *Response, error) {
	url := fmt.Sprintf("masters/%d", id)

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	master := new(Master)
	resp, err := s.client.Do(req, master)

	if err != nil {
		return nil, resp, err
	}

	return master, resp, err
}

func (s *MasterService) GetVersions(id int) (*MasterVersions, *Response, error) {
	url := fmt.Sprintf("masters/%d/versions", id)

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	masterVersions := new(MasterVersions)
	resp, err := s.client.Do(req, masterVersions)

	if err != nil {
		return nil, resp, err
	}

	return masterVersions, resp, err
}
