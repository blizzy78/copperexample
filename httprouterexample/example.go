package httprouterexample

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/blizzy78/copper/template"

	"github.com/blizzy78/copperexample/middleware"
)

func Run(r *template.Renderer) {
	router := httprouter.New()

	h := index(r)
	h = middleware.NewRequestID(h)

	router.Handler("GET", "/", h)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(renderer *template.Renderer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		data := map[string]interface{}{
			"title": "httprouter",
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
