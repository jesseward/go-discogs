package discogs

type Community struct {
	Status       string         `json:"status"`
	Rating       Rating         `json:"rating"`
	Want         int            `json:"want"`
	Contributors []UserResource `json:"contributors"`
	Have         int            `json:"have"`
	Submitter    UserResource   `json:"submitter"`
	DataQuality  string         `json:"data_quality"`
}

type Rating struct {
	Count   int     `json:"count"`
	Average float64 `json:"average"`
}

type UserResource struct {
	Username    string `json:"username"`
	ResourceURL string `json:"resource_url"`
}
