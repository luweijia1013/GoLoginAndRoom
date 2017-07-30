package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Arith struct {
	client *rpc.Client
}

func (t *Arith) Divide(a, b int) Quotient {
	args := &Args{a, b}
	var reply Quotient
	err := t.client.Call("Arithmetic.Divide", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	return reply
}

func (t *Arith) Multiply(a, b int) int {
	args := &Args{a, b}
	var reply int
	err := t.client.Call("Arithmetic.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	return reply
}

func main() {

	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Connectiong:", err)
	}

	arith := &Arith{client: rpc.NewClient(conn)}

	fmt.Println(arith.Multiply(5, 6))
	fmt.Println(arith.Divide(500, 10))
}
