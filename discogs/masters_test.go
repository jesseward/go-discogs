package discogs

import (
	"fmt"
	"testing"
)

// TestGetMaster fetches Master metadata based on the master id
func TestGetMaster(t *testing.T) {

	// Master Title we expect for master id = 819704
	expectedResult := "DJ-Kicks"

	client := getClient(t)
	master, _, err := client.Database.GetMaster(819704)
	check(t, err)

	assert(t, master.Title == expectedResult, fmt.Sprintf("Master.Title expected %s, received %s ", expectedResult, master.Title))
}

// TestMasterGetVersions fetches Master Release Versions based on the master id
func TestGetMasterVersions(t *testing.T) {

	// We expect Pagination.pages >= 1 for master is 819704
	expectedResult := 1

	client := getClient(t)
	versions, _, err := client.Database.GetMasterVersions(819704)
	check(t, err)

	assert(t, versions.Pagination.Pages >= expectedResult, fmt.Sprintf("Master.Pagination.Pages expected %d, received %d", expectedResult, versions.Pagination.Pages))

}
