package main

import (
	todo_router "gt/chap31/ex/todo/router"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	mux.Use(MimeContentTypeJSONMiddleware)
	todo_router.Load(mux)
	http.ListenAndServe(":3250", handlers.CORS(
		handlers.AllowedHeaders(
			[]string {
				"X-Requested-With", "Content-Type", "Authorization", "Content-Length", "X-CSRF-Token", "Accept", "HRVGameLocation", "X-Frame-Options",
			},
		),
		handlers.AllowedMethods(
			[]string {
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodDelete,
			},
		),
	)(mux))
}
func MimeContentTypeJSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
