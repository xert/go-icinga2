package icinga

import (
	"net/http"
	"testing"
)

func TestBasicAuthTransport(t *testing.T) {
	setup()
	defer teardown()

	username, password := "u", "p"

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok {
			t.Errorf("request does not contain basic auth credentials")
		}
		if u != username {
			t.Errorf("request contained basic auth username %q, want %q", u, username)
		}
		if p != password {
			t.Errorf("request contained basic auth password %q, want %q", p, password)
		}
	})

	tp := &BasicAuthTransport{
		Username: username,
		Password: password,
	}
	basicAuthClient := NewClient(baseURL, tp.Client())
	basicAuthClient.BaseURL = client.BaseURL
	req, _ := basicAuthClient.NewRequest("GET", "/", nil)
	_, err := basicAuthClient.Do(req, nil)
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}
}

func TestBasicAuthTransport_transport(t *testing.T) {
	// default transport
	tp := &BasicAuthTransport{}
	if tp.transport() != http.DefaultTransport {
		t.Errorf("Expected http.DefaultTransport to be used.")
	}

	// custom transport
	tp = &BasicAuthTransport{
		Transport: &http.Transport{},
	}
	if tp.transport() == http.DefaultTransport {
		t.Errorf("Expected custom transport to be used.")
	}
}
