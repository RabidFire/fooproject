package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/RabidFire/fooproject/pkg/api"
)

func TestRootHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		panic(err)
	}

	rr := httptest.NewRecorder()

	http.HandlerFunc(RootHandler).ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Error("expected status 200, got:", rr.Code)
	}

	if rr.Body.String() != "Hello World" {
		t.Error("expected body Hello World, got:", rr.Body.String())
	}
}
