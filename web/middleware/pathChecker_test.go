package middleware

import "testing"

func TestBannish(t *testing.T) {
	url := `host/GetGames/$insert`
	expected := true
	if expected != bannish(url) {
		t.Errorf("bannish func isn't working")
	}
}
