package discogs

import "testing"

// Check an error result for a value.  If present, fail the test with an error written to the console.
func check(t *testing.T, e error) {
	if e != nil {
		t.Error(e)
	}
}

// assert provides a simple way to verify an assertion, and fail the test if that assertion fails.
func assert(t *testing.T, condition bool, assertion string) {
	if !condition {
		t.Errorf("Assertion failed: %v", assertion)
	}
}

// creates a testing HTTP client.
// TODO: unit tests should be *mocked* . May be helpful -> http://keighl.com/post/mocking-http-responses-in-golang/
func getClient(t *testing.T) *Client {
	client := NewClient(nil)
	return client
}
