package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

var httpAddr = flag.String("http", ":8080", "Listen address")

func main() {
	flag.Parse() //fra kommandolinje til kjørring
	server := NewServer()
	log.Fatal(http.ListenAndServe(*httpAddr, server))
}

// Server implements the web server specification found at
// lab2/README.md#web-server
type Server struct {
	mux    *http.ServeMux
	rcount int
}

// NewServer returns a new Server with all required internal state initialized.
// NOTE: It should NOT start to listen on an HTTP endpoint.
func NewServer() *Server {
	s := &Server{
		mux: http.NewServeMux(),
	}
	// TODO(student): Implement
	s.mux.HandleFunc("/", s.root)
	s.mux.HandleFunc("/counter", s.counter)
	s.mux.HandleFunc("/lab2", s.lab2)
	s.mux.HandleFunc("/fizzbuzz", s.fizzbuzz)
	return s
}
func (s *Server) root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		writeLogErr(w, "404 page not found\n")
		return
	}
	writeLogErr(w, "Hello World!\n")
}

func (s *Server) counter(w http.ResponseWriter, r *http.Request) {
	writeLogErr(w, fmt.Sprintf("counter: %v\n", s.rcount))
}

func (s *Server) fizzbuzz(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("value")
	if value == "" {
		writeLogErr(w, "no value provided\n")
		return
	}
	num, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		writeLogErr(w, "not an integer\n")
		return
	}

	var result string
	if num%3 == 0 {
		result += "fizz"
	}
	if num%5 == 0 {
		result += "buzz"
	}
	if result == "" {
		result = strconv.FormatInt(num, 10)
	}
	result += "\n"
	writeLogErr(w, result)
}

func (s *Server) lab2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMovedPermanently)
	writeLogErr(w, "<a href=\"http://www.github.com/uis-dat520/labs/tree/master/lab2\">Moved Permanently</a>.\n\n")
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO(student): Implement
	s.rcount++
	s.mux.ServeHTTP(w, r)
}

func writeLogErr(wr io.Writer, s string) {
	_, err := fmt.Fprint(wr, s)
	if err != nil {
		log.Println(err)
	}

}
