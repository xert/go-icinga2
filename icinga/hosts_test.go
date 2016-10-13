package icinga

import (
	"encoding/json"
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
		fmt.Fprint(w, `{ "results": [ { "attrs": { "__name": "1" } } ] }`)
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

func TestHostsService_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		w.WriteHeader(http.StatusOK)
	})

	_, res, _ := client.Hosts.Create(&Host{Name: "A", CheckCommand: "hostalive", Address: "10.0.0.1"})

	if res == nil {
		t.Fatal("Response is nil")
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d got %d", http.StatusOK, res.StatusCode)
	}
}

func TestHostsService_CreateBody(t *testing.T) {
	setup()
	defer teardown()

	host := Host{Address: "localhost", Name: "A"}
	wantHost := hostAttrs{
		Attrs: host,
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, "Accept", mediaTypeJSON)
		v := new(hostAttrs)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(wantHost, *v) {
			t.Errorf("Expected %v got %v", wantHost, v)
		}
	})

	_, res, _ := client.Hosts.Create(&Host{
		Name:    host.Name,
		Address: host.Address,
	})

	if res == nil {
		t.Fatal("Response is nil")
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d got %d", http.StatusOK, res.StatusCode)
	}
}
