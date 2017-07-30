package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"os"
	"net/http"
	"io/ioutil"
)

var (
	addr = flag.String("addr", "localhost:8000", "tcp dial address")
)

func LoginJudge(res http.ResponseWriter, req *http.Request) {
	//user :=
}

func main() {
	flag.Parse()
	log.Println("start robot")

	/*
	**  loginState: 0->invalid, 1->valid but not send, 2->sended
	 */
	var loginState = 0
	var name string = ""
	var pass string = ""
	var loginResult string = ""
	stdin_login := bufio.NewReader(os.Stdin)
	for {
		line, _, _ := stdin_login.ReadLine()
		if name == "" {
			name = string(line)
		} else if name != "" && pass == "" {
			pass = string(line)
			if pass != ""{
				break
			}
		}
	}

	response, _ := http.Get("http://localhost:8888/login?user=" + name + "&pass=" + pass)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	loginResult = string(body)

	if loginResult == "OK" {
		log.Println("login sccuess")
		loginState = 1;
	}else{
		log.Println("login fail")
		return
	}

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
		if (loginState == 2) {
			conn.Write([]byte(string(line) + "\n"))
		} else if (loginState == 1) {
			conn.Write([]byte("[login]\n" + string(line) + "\n"))
			loginState = 2
		}
	}
}
