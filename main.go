package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

var appHTML string
var indexHTML string
var imagesList []string

func main() {
	appHTMLBinary, err := os.ReadFile("src/app.html")
	if err != nil {
		panic(err)
	}
	appHTML = string(appHTMLBinary)

	indexHTMLBinary, err := os.ReadFile("src/index.html")
	if err != nil {
		panic(err)
	}
	indexHTML = string(indexHTMLBinary)

	files, err := os.ReadDir("static/images")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		imagesList = append(imagesList, file.Name())
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") && r.URL.Path != "/" {
			http.Redirect(w, r, strings.TrimSuffix(r.URL.Path, "/"), http.StatusMovedPermanently)

		} else if strings.HasPrefix(r.URL.Path, "/static") {
			http.ServeFile(w, r, "."+r.URL.Path)

		} else if r.URL.Path == "/" {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(indexHandler()))
		}
	})

	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func indexHandler() string {
	pageHTML := strings.ReplaceAll(appHTML, "{%app%}", indexHTML)
	pageHTML = strings.ReplaceAll(pageHTML, "{%backgroundImage%}", fmt.Sprintf("/static/images/%s", randomImage()))

	return pageHTML
}

func randomImage() string {
	randomIndex := rand.Intn(len(imagesList))
	return imagesList[randomIndex]
}
