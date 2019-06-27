package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type dumbAuth struct {
	trigger string
}

func (a *dumbAuth) IsAuthorized(s string) bool {
	return s != "" && s != a.trigger
}

func alwaysGood(w http.ResponseWriter, r *http.Request) {}

func TestAuthorizeFunc(t *testing.T) {
	tests := map[string]struct {
		params   string
		wantCode int
	}{
		"basic, valid":         {"auth=junk", http.StatusOK},
		"basic, invalid":       {"aut=junk", http.StatusForbidden},
		"authorizer is called": {"auth=trigger", http.StatusForbidden},
	}

	for tn, tt := range tests {
		req := httptest.NewRequest("GET", "http://example.com/foo?"+tt.params, nil)
		w := httptest.NewRecorder()
		authorize := authorizeFunc("auth", &dumbAuth{"trigger"})
		authorize(http.HandlerFunc(alwaysGood)).ServeHTTP(w, req)

		res := w.Result()
		res.Body.Close()

		got := res.StatusCode
		if got != tt.wantCode {
			t.Errorf("%s: got %v, want %v", tn, got, tt.wantCode)
		}
	}
}
