package discogs

import (
	"fmt"
	"testing"
)

// TestGetMaster fetches Master metadata based on the master id
func TestMasterGet(t *testing.T) {

	// Master Title we expect for master id = 819704
	expectedResult := "DJ-Kicks"

	client := getClient(t)
	master, _, err := client.Master.Get(819704)
	check(t, err)

	assert(t, master.Title == expectedResult, fmt.Sprintf("Master.Title expected %s, received %s ", expectedResult, master.Title))
}

// TestMasterGetVersions fetches Master Release Versions based on the master id
func TestMasterGetVersions(t *testing.T) {

	// We expect Pagination.pages >= 1 for master is 819704
	expectedResult := 1

	client := getClient(t)
	versions, _, err := client.Master.GetVersions(819704)
	check(t, err)

	assert(t, versions.Pagination.Pages >= expectedResult, fmt.Sprintf("Master.Pagination.Pages expected %d, received %d", expectedResult, versions.Pagination.Pages))

}
