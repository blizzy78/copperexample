package main

import (
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
	r := template.NewRenderer(template.LoaderFunc(loadTemplate),
		template.WithScopeData("safe", helpers.Safe),
		template.WithScopeData("html", helpers.HTML),
		template.WithScopeData("has", helpers.Has),
		template.WithScopeData("requestID", middleware.RequestIDFromContext),
	)

	// chiexample.Run(r)
	// gorillaexample.Run(r)
	// httprouterexample.Run(r)
	nethttpexample.Run(r)
}

func loadTemplate(name string) (io.ReadCloser, error) {
	path := filepath.FromSlash("templates" + name + ".html")
	return os.Open(path)
}
