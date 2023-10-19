package assertions

import "testing"

func ExpectToBe(t *testing.T, value interface{}, expectedValue interface{}) {
	errorMessage := "\nExpected: %v\nReceived: %v"

	if value != expectedValue {
		t.Errorf(errorMessage, expectedValue, value)
	}
}
