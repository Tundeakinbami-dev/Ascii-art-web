package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)


var tmpl = template.Must(template.ParseFiles("index.html"))

func main() {

    mux := http.NewServeMux()
    mux.HandleFunc("/home", homeHandler)
    mux.HandleFunc("/ascii-art", asciiHandler)

    fmt.Println("The server is running on: http://localhost:7071")

    log.Fatal(http.ListenAndServe(":7071", mux))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/home" {
        http.Error(w, "404 Not Found", http.StatusNotFound)
        return
    }
    
    err := tmpl.Execute(w, nil) 
    if err != nil {
        http.Error(w, "500 Internal server Error", http.StatusInternalServerError)
    }
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "400 Bad Request", http.StatusBadRequest)
        return
    }

   
    text := r.FormValue("textname")
    banner := r.FormValue("banner")

 

    bannerMap, err := loadBanner("banners/" + banner + ".txt") 
    if err != nil {
        http.Error(w, "404 banner not found", http.StatusNotFound)
    }

    splitText := strings.Split(text, "\n")
    result := render(splitText, bannerMap)
    
    err = tmpl.Execute(w, result) 
    if err != nil {
        http.Error(w, "500 Internal server Error", http.StatusInternalServerError)
    }
}


