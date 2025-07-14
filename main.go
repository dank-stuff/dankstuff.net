package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
)

var (
	//go:embed templates
	aaa embed.FS

	//go:embed assets
	bbb embed.FS
)

func main() {
	t := template.Must(template.ParseFS(aaa, "templates/index.gohtml"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = t.Execute(w, map[string]any{
			"BG": fmt.Sprintf("/assets/bgs/%d.webp", rand.Intn(5)+1),
			"DankStuffProducts": []struct {
				Title       string
				Link        string
				Description string
				LogoPath    string
				WIP         bool
			}{

				{Title: "DankMuzikk", Link: "https://dankmuzikk.com", Description: "Create, Share and Play Music Playlists.", LogoPath: "/assets/logos/dankmuzikk.webp"},
				{Title: "DankLyrics", Link: "https://danklyrics.com", Description: "Find lyrics for songs or something.", LogoPath: "/assets/logos/danklyrics.png"},
				{Title: "DankScreen", Link: "https://screen.dankstuff.net", Description: "Display capture card's output into your browser.", LogoPath: "/assets/logos/dankscreen.png"},
				{Title: "DankTodo", Link: "https://todo.dankstuff.net", Description: "The first htmx app with C (Ulfius)", LogoPath: "/assets/logos/danktodo.png"},
				{Title: "DankNotes", Link: "https://notes.dankstuff.net", Description: "Bootleg version of Notion and Google Notes.", WIP: true},
				// {Title: "DankSim", Link: "https://danksim.com", Description: "Get an eSIM qucickly, ANYWHERE.", WIP: true},
				{Title: "DankDysk", Link: "https://dankdysk.com", Description: "Magical unlimited cloud drive.", WIP: true},
				{Title: "DankTorrent", Link: "https://dankdysk.com", Description: "Download anything as torrent.", WIP: true},
				{Title: "DankIP", Link: "https://dankip.com", Description: "Get a public IP for any server/computer.", WIP: true},
				{Title: "libdank", Link: "https://libdank.org", Description: "File compression extension but the result is a video.", WIP: true},
			},
		})
	})

	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		robotsFile, _ := bbb.ReadFile("assets/robots.txt")
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write(robotsFile)
	})
	http.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		robotsFile, _ := bbb.ReadFile("assets/sitemap.xml")
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write(robotsFile)
	})

	http.Handle("/assets/", http.FileServer(http.FS(bbb)))

	log.Println("server running on port 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
