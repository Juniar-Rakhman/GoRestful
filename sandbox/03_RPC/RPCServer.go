package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct {
	Foo, Bar string
}

type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	log.Printf("foo: %s", args.Foo)
	log.Printf("bar: %s", args.Bar)
	// Fill reply pointer to send the data back
	*reply = time.Now().Unix()
	return nil
}

func main() {

	// Create a new RPC server
	timeserver := new(TimeServer)

	// Register RPC server
	rpc.Register(timeserver)
	rpc.HandleHTTP()

	// Listen for requests on port 1234
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(l, nil)
}
