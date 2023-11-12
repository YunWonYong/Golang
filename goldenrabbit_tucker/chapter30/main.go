package main

import (
	student_router "gt/chap30/ex/student/router"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	mux.Use(MimeContentTypeJSONMiddleware)
	student_router.Load(mux)
	http.ListenAndServe(":3000", mux)
}
func MimeContentTypeJSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
