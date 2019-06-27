package main

import "net/http"

type authorizer interface {
	IsAuthorized(string) bool
}

func authorizeFunc(authKey string, a authorizer) func(http.Handler) http.Handler {

	// authKey is valid here
	// carrying this scope everywhere "authorize" goes

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authVals, ok := r.URL.Query()[authKey]
			if !ok || !a.IsAuthorized(authVals[0]) {
				stts := http.StatusForbidden
				http.Error(w, http.StatusText(stts), stts)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
