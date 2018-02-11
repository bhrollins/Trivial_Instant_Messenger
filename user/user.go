package user

import (
  "fmt"
  "net"
)

type User struct {
  password string
  username string
  queuedMessages []string
  conn net.Conn
}

// constructor
func NewUser(n, p string) *User {
  u := new(User)
  u.password = p
  u.username = n
  u.queuedMessages = []string{}

  return u
}

// --------------------------- getters ---------------------------------------
func (usr User) Username() string {
  return usr.username
}

func (usr User) Messages() []string {
  return usr.queuedMessages
}

// -------------------- pointer-receiver functions ----------------------------
func (usr *User) Authenticate(n, p string) bool {
  if usr.username == n && usr.password == p {
    return true
  }

  return false
}

func (usr *User) ClearMessages() {
  usr.queuedMessages = []string{}
}

func (usr *User) Connect(c net.Conn) {
  usr.conn = c
}

func (usr *User) Disconnect() {
  usr.conn = nil
}

// ---------------------- value-receiver functions ----------------------------
func (usr User) Connected() bool {
  return usr.conn == nil
}

func (usr User) Send(from, mess string) {
  fmt.Fprintf(usr.conn, "Message from %s as follows: %s\r\n", from, mess)
}

func (usr User) GetConn() net.Conn {
  return usr.conn
}
