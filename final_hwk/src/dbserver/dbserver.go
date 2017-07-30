package main

import (
	"fmt"
	"net/http"
	"net/rpc"
	"strings"
)

var S string

type MyRPC int

func (r *MyRPC) LoginTest(S string, reply *string) error {
	fmt.Println(S)
	s := strings.Split(S, ",")
	if len(s) == 2 && s[0] == s[1] {
		*reply = "OK"
	}else{
		*reply = "FAIL"
	}
	return nil
}


func main() {
	fmt.Println("start dbserver")
	r := new(MyRPC)
	rpc.Register(r)
	rpc.HandleHTTP()
	err := http.ListenAndServe("localhost:7777", nil)
	if err != nil {
		fmt.Println("rpc connection", err.Error())
	}
}
