package httprouterexample

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/blizzy78/copper/template"

	"github.com/blizzy78/copperexample/middleware"
)

func Run(rd *template.Renderer) {
	router := httprouter.New()

	h := index(rd)
	h = middleware.NewRequestID(h)

	router.Handler("GET", "/", h)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(rd *template.Renderer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"title": "httprouter",
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
