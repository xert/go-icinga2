package icinga

import "testing"

func TestNewClient(t *testing.T) {
	baseURL := "https://localhost:5665/"
	c := NewClient(baseURL, nil)

	if got, want := c.BaseURL.String(), baseURL+apiVersion+"/"; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}
	if got, want := c.UserAgent, userAgent; got != want {
		t.Errorf("NewClient UserAgent is %v, want %v", got, want)
	}
}
