package main

import (
	"errors"
	"log"
	"net"
)

var (
	errUser = errors.New("error user")
)

func main() {
	log.Println("start room server")

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Println("listen error", err)
		return
	}
	var id uint32
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("accept error ", err)
		}
		id++
		go NewUser(conn, id).Loop()
	}
}
