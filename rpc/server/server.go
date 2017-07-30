package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
)

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	log.Println("call by client")
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	log.Println("call by client")
	return nil
}

func main() {

	arith := new(Arith)

	server := rpc.NewServer()
	server.RegisterName("Arithmetic", arith)

	addr := ":1234"
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	log.Println("listen port ", addr)

	server.Accept(l)
}
