package main

import (
	"errors"
	"log"
	"flag"
	"net/http"
	"net/rpc"
)

var (
	errUser = errors.New("error user")
)

func LoginJudge(res http.ResponseWriter, req *http.Request){
	user := req.FormValue("user")
	pass := req.FormValue("pass")

	client, err := rpc.DialHTTP("tcp", "localhost:7777")
	if err != nil {
		log.Fatal("dialhttp: ", err)
	}

	var reply *string
	S := user+","+pass
	err = client.Call("MyRPC.LoginTest", S, &reply)
	if err != nil {
		log.Fatal("call rpc: ", err)
	}
	log.Println("REPLY:",*reply)
	if *reply == "OK"  {
		res.Write([]byte("OK"))
		log.Println("login success")
	}else{
		res.Write([]byte("FAIL"))
		log.Println("login fail")
	}
}

func main() {
	log.Println("start login server")

	host := flag.String("host", "127.0.0.1", "listen host")
	port := flag.String("port", "8888", "listen port")
	http.HandleFunc("/login", LoginJudge)
	err := http.ListenAndServe(*host + ":" + *port, nil)
	if err != nil {
		log.Println("login error!", err)
		return
	}
}
