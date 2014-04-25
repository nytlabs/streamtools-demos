package main

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(body)
}

func TutorialHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	body, err := ioutil.ReadFile(vars["tutorial"] + ".html")
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(body)
}

func CssHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	body, err := ioutil.ReadFile("css/" + vars["filename"])
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(body)
}

func JsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	body, err := ioutil.ReadFile("js/" + vars["filename"])
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(body)
}

func main() {
	log.Println("streamtools demos are running on http://localhost:8080/")
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)
	r.HandleFunc("/{tutorial}.html", TutorialHandler)
	r.HandleFunc("/css/{filename}", CssHandler)
	r.HandleFunc("/js/{filename}", JsHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
