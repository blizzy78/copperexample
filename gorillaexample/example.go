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

func Run(rd *template.Renderer) {
	m := mux.NewRouter()
	m.Use(handlers.RecoveryHandler())

	h := index(rd)
	h = middleware.NewRequestID(h)

	m.Handle("/", h)

	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, m)))
}

func index(rd *template.Renderer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"title": "Gorilla",
			"user": map[string]interface{}{
				"firstName": "John",
				"lastName":  "Doe",
				"age":       42,
			},
		}

		if err := rd.Render(r.Context(), w, "/index", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
