package nethttpexample

import (
	"log"
	"net/http"

	"github.com/blizzy78/copper/template"

	"github.com/blizzy78/copperexample/middleware"
)

func Run(r *template.Renderer) {
	var h http.Handler
	h = index(r)
	h = middleware.NewRequestID(h)

	http.Handle("/", h)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(renderer *template.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		data := map[string]interface{}{
			"title": "net/http",
			"user": map[string]interface{}{
				"firstName": "John",
				"lastName":  "Doe",
				"age":       42,
			},
		}

		if err := renderer.Render(ctx, w, "/index", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
