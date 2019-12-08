package gorillaexample

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/blizzy78/copper/template"

	"github.com/blizzy78/copperexample/middleware"
)

func Run(r *template.Renderer) {
	m := mux.NewRouter()
	m.Use(handlers.RecoveryHandler())

	h := index(r)
	h = middleware.NewRequestID(h)

	m.Handle("/", h)

	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, m)))
}

func index(renderer *template.Renderer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		data := map[string]interface{}{
			"title": "Gorilla",
			"user": map[string]interface{}{
				"firstName": "John",
				"lastName":  "Doe",
				"age":       42,
			},
		}

		if err := renderer.Render(ctx, w, "/index", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
