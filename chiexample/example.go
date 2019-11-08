package chiexample

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/blizzy78/copper/template"

	exmiddleware "github.com/blizzy78/copperexample/middleware"
)

func Run(r *template.Renderer) {
	router := chi.NewRouter()
	router.Use(middleware.Logger, middleware.Recoverer, exmiddleware.NewRequestID)

	router.Get("/", index(r))

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(renderer *template.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		data := map[string]interface{}{
			"title": "chi",
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
