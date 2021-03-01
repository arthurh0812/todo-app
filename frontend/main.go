package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var port = flag.String("port", "8000", "the port to listen on")
var host = flag.String("host", "127.0.0.1", "the host name to listen on")

type FileSystem struct {
	fs http.FileSystem
}

func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}

func main() {
	// parse flags
	flag.Parse()

	mux := http.NewServeMux()

	staticFileServer := http.FileServer(FileSystem{http.Dir("./static")})

	mux.Handle("/static/", http.StripPrefix("/static", staticFileServer))

	base, err := template.New("base.html").ParseFiles("./templates/base.html")
	data := Page{
		Title: "TODO APP",
		Header: "nice todo app!",
	}

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Hello Docker!")

	err = base.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello Docker!"))
		err := base.Execute(w, data)
		if err != nil {
			log.Fatal(err)
		}
	})

	srv := &http.Server{
		Addr: *host + ":" + *port,
		Handler: mux,
	}

	log.Printf("Listening to requests at %s", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}
