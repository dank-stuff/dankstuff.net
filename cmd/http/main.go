package main

import (
	"context"
	"log"
	"net/http"
	"regexp"

	"dankstuff.net/assets"
	"dankstuff.net/layouts"
	"dankstuff.net/pages"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	mjson "github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
	"github.com/tdewolff/minify/v2/xml"
)

func main() {
	minifyer := minify.New()
	minifyer.AddFunc("text/css", css.Minify)
	minifyer.AddFunc("text/html", html.Minify)
	minifyer.AddFunc("image/svg+xml", svg.Minify)
	minifyer.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	minifyer.AddFuncRegexp(regexp.MustCompile("[/+]json$"), mjson.Minify)
	minifyer.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)

	assets := assets.FS()
	pagesHandler := http.NewServeMux()

	pagesHandler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "path-name", "/")
		layouts.Default(pages.Home()).Render(ctx, w)
	})

	pagesHandler.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "path-name", "/products")
		layouts.Default(pages.Products()).Render(ctx, w)
	})

	pagesHandler.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "path-name", "/about")
		layouts.Default(pages.About()).Render(ctx, w)
	})

	pagesHandler.HandleFunc("/privacy", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "path-name", "/privacy")
		layouts.Default(pages.Privacy()).Render(ctx, w)
	})

	pagesHandler.HandleFunc("/terms-of-use", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "path-name", "/terms-of-use")
		layouts.Default(pages.TermsOfUse()).Render(ctx, w)
	})

	handler := http.NewServeMux()
	handler.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		robotsFile, _ := assets.ReadFile("robots.txt")
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write(robotsFile)
	})
	handler.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		sitemapFile, _ := assets.ReadFile("sitemap.xml")
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write(sitemapFile)
	})

	handler.Handle("/assets/", minifyer.Middleware(http.StripPrefix("/assets", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=7200, stale-while-revalidate=5")
		http.FileServer(http.FS(assets)).ServeHTTP(w, r)
	}))))

	handler.Handle("/", minifyer.Middleware(pagesHandler))

	log.Println("server running on port 8080")
	log.Fatalln(http.ListenAndServe(":8080", handler))
}
