package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/blizzy78/copper/helpers"
	"github.com/blizzy78/copper/template"

	_ "github.com/blizzy78/copperexample/chiexample"
	_ "github.com/blizzy78/copperexample/gorillaexample"
	_ "github.com/blizzy78/copperexample/httprouterexample"
	"github.com/blizzy78/copperexample/nethttpexample"

	"github.com/blizzy78/copperexample/middleware"
)

func main() {
	scopeData := map[string]interface{}{
		"safe":      helpers.Safe,
		"html":      helpers.HTML,
		"has":       helpers.Has,
		"requestID": middleware.RequestIDFromContext,
	}

	r := template.NewRenderer(loadTemplate, nil, scopeData, "t")

	nethttpexample.Run(r)
	// gorillaexample.Run(r)
	// chiexample.Run(r)
	// httprouterexample.Run(r)
}

func loadTemplate(name string) (io.Reader, error) {
	path := fmt.Sprintf("templates%s.html", name)
	path = filepath.FromSlash(path)
	return os.Open(path)
}
