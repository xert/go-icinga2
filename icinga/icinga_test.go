package icinga

import (
	"io/ioutil"
	"testing"
)

const (
	baseURL = "https://localhost:5665/"
)

func TestNewClient(t *testing.T) {
	c := NewClient(baseURL, nil)

	if got, want := c.BaseURL.String(), baseURL+apiVersion+"/"; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}
	if got, want := c.UserAgent, userAgent; got != want {
		t.Errorf("NewClient UserAgent is %v, want %v", got, want)
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient(baseURL, nil)

	inURL, outURL := "/foo", baseURL+"foo"
	inBody, outBody := &Host{Name: "u"}, `{"__name":"u"}`+"\n"
	req, _ := c.NewRequest("GET", inURL, inBody)

	// test that relative URL was expanded
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}

	// test that body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("NewRequest(q) Body is %v, want %v", got, want)
	}

	// test that default user-agent is attached to the request
	if got, want := req.Header.Get("User-Agent"), c.UserAgent; got != want {
		t.Errorf("NewRequest() User-Agent is %v, want %v", got, want)
	}
}
