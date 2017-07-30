package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"os"
)

var (
	addr = flag.String("addr", "localhost:8000", "tcp dial address")
)

func main() {
	flag.Parse()
	log.Println("start robot")

	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		log.Println("connect error!", err)
		return
	}

	reader := bufio.NewReader(conn)

	go func() {
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				log.Println("read error!", err)
				return
			}
			log.Println("recv data", string(line))
		}
	}()

	stdin := bufio.NewReader(os.Stdin)
	for {
		line, _, _ := stdin.ReadLine()
		conn.Write([]byte(string(line) + "\n"))
	}
}
