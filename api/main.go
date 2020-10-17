package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) routes() {
	s.router.HandleFunc("/", s.handleEmojis())
	s.router.HandleFunc("/_healthcheck", s.handleHealthCheck())
}

func (s *server) handleEmojis() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		endpoint := fmt.Sprintf("http://tracker.%s:3000/", os.Getenv("COPILOT_SERVICE_DISCOVERY_ENDPOINT"))
		resp, err := http.Get(endpoint)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}

func (s *server) handleHealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("healthcheck okay!"))
	}
}

func main() {
	addr := flag.String("addr", ":8080", "port to listen on")
	flag.Parse()
	log.Printf("port %s\n", *addr)

	handler := &server{
		router: mux.NewRouter(),
	}
	handler.routes()

	s := &http.Server{
		Addr:         *addr,
		Handler:      handler,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
	log.Fatal(s.ListenAndServe())
}
