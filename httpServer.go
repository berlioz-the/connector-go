package berlioz

import (
	"net/http"
)

// TBD
func WrapFunc(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		newReq, span := myZipkin.instrumentServerRequest(req)
		defer span.Finish()
		h(w, newReq)
	})
}
