package main

import (
	"bufio"
	"log"
	"net"
	"strings"
	"sync"
	"sync/atomic"
)

type User struct {
	Id      uint32
	Name    string
	conn    net.Conn
	isClose int32
	isLogin bool
}

func (this *User) IsClosed() bool {
	return atomic.LoadInt32(&this.isClose) == 0
}

func (this *User) Loop() {
	reader := bufio.NewReader(this.conn)

	for !this.IsClosed() {
		line, _, err := reader.ReadLine()
		if err != nil {
			this.Close()
			return
		}

		if strings.HasPrefix(string(line), "[login]") {
			this.isLogin = true
			continue
		}
		if !this.isLogin {
			this.write("please login")
			continue
		}

		log.Println("recv data:", string(line), ", from ", this.conn.RemoteAddr())

		UserMgr_GetMe().SendAll(string(line), this.Id)
	}
}

func (this *User) write(str string) (err error) {
	_, err = this.conn.Write([]byte(str + "\n"))
	return
}

func (this *User) Close() {
	if atomic.CompareAndSwapInt32(&this.isClose, 1, 0) {
		// TODO Close
		if this.conn != nil {
			this.conn.Close()
		}
	}
}

func NewUser(conn net.Conn, id uint32) (u *User) {
	u = &User{
		conn:    conn,
		isClose: 1,
		Id:      id,
	}
	UserMgr_GetMe().Add(u)
	return u
}

var _userMsg *UserMgr

func UserMgr_GetMe() *UserMgr {
	if _userMsg == nil {
		_userMsg = &UserMgr{
			users: make(map[uint32]*User),
		}
	}
	return _userMsg
}

type UserMgr struct {
	users map[uint32]*User
	sync.Mutex
}

func (this *UserMgr) Add(u *User) {
	this.Lock()
	this.users[u.Id] = u
	this.Unlock()
}

func (this *UserMgr) Del(u *User) {
	this.Lock()
	delete(this.users, u.Id)
	this.Unlock()
}

func (this *UserMgr) SendAll(msg string, eid uint32) {
	this.Lock()
	for _, u := range this.users {
		if u.Id == eid {
			continue
		}
		u.write(msg)
	}
	this.Unlock()
}
