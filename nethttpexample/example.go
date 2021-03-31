package nethttpexample

import (
	"log"
	"net/http"

	"github.com/blizzy78/copper/template"

	"github.com/blizzy78/copperexample/middleware"
)

func Run(rd *template.Renderer) {
	h := index(rd)
	h = middleware.NewRequestID(h)

	http.Handle("/", h)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(rd *template.Renderer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"title": "net/http",
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
