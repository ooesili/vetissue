package vetissue

import "testing"

func TestValue(t *testing.T) {
	if !value {
		t.Errorf("value should be true, got false")
	}
}
