package icinga

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestHostsService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+apiVersion+"/objects/hosts/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", mediaTypeJSON)
		fmt.Fprint(w, `{ "results": [ { "attrs": { "__name": "1" } } ] } `)
	})

	host, _, err := client.Hosts.Get("1")
	if err != nil {
		t.Errorf("Hosts.Get returned error: %v", err)
	}

	want := &Host{
		Name: "1",
	}
	if !reflect.DeepEqual(host, want) {
		t.Errorf("Hosts.Get returned %+v, want %+v", host, want)
	}
}

func TestHostsService_Get_DoError(t *testing.T) {
	setup()
	defer teardown()

	_, _, err := client.Hosts.Get("1")

	if err != nil {
		e := err.(*ErrorResponse)
		want := 404
		if e.Response.StatusCode != want {
			t.Errorf("Hosts.Get returned %+v, want %+v", e.Response.StatusCode, want)
		}
	} else {
		t.Error("Expected HTTP error.")
	}
}
