package discogs

import (
	"fmt"
	"testing"
)

// TestGetLabel fetches Label metadata based on the label id
func TestLabelGet(t *testing.T) {

	// Label Name we expect for label id = 23528
	expectedResult := "Warp Records"

	client := getClient(t)
	label, _, err := client.Label.Get(23528)
	check(t, err)

	assert(t, label.Name == expectedResult, fmt.Sprintf("Label.Name expected %s, received %s ", expectedResult, label.Name))
}

// TestLabelGetReleases fetches Label Releases based on the label id
func TestLabelGetReleases(t *testing.T) {

	// We expect Pagination.pages == 60 for label is 23528
	expectedResult := 60

	client := getClient(t)
	releases, _, err := client.Label.GetReleases(23528)
	check(t, err)

	assert(t, releases.Pagination.Pages >= expectedResult, fmt.Sprintf("Label.Pagination.Pages expected %d, received %d", expectedResult, releases.Pagination.Pages))

}
