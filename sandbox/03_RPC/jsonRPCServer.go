package main

import (
	jsonparse "encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Args holds arguments passed to JSON RPC service
type Args struct {
	Id string
}

// Book struct holds Book JSON structure
type Book struct {
	Id     string `json:"id, omitempty"`
	Name   string `json:"name, omitempty"`
	Author string `json:"author, omitempty"`
}

type JSONServer struct{}

func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book

	// Read JSON file and load data
	raw, err := ioutil.ReadFile("/home/jrakhman/go/src/github.com/Juniar-Rakhman/GoRestful/sandbox/03_RPC/books.json")
	if err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}

	// Unmarshal JSON raw data into books array
	if err := jsonparse.Unmarshal(raw, &books); err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}

	// Iterate over each book to find the given book
	for _, book := range books {
		if book.Id == args.Id {
			// If book found, fill reply with it
			*reply = book
			break
		}
	}
	return nil
}
func main() {
	// Create a new RPC server
	s := rpc.NewServer() // Register the type of data requested as JSON
	s.RegisterCodec(json.NewCodec(), "application/json")

	// Register the service by creating a new JSON server
	s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":1234", r)
}
