package main

import (
	"log"
	"net/rpc"
)

type Args struct {
	Foo, Bar string
}

func main() {
	var reply int64
	args := Args{"foo", "bar"}

	client, err := rpc.DialHTTP("tcp", "localhost"+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	err = client.Call("TimeServer.GiveServerTime", args, &reply)
	if err != nil {
		log.Fatal("Error:", err)
	}

	log.Printf("%d", reply)
}
