package user

import (
  "fmt"
  "net"
)

type User struct {
  password string
  username string
  queuedMessages map[string][]string
  conn net.Conn
  isConnected bool
}

// constructor
func NewUser(n, p string) *User {
  u := new(User)
  u.password = p
  u.username = n
  u.queuedMessages = make(map[string][]string)
  u.isConnected = false

  return u
}

// --------------------------- getters ---------------------------------------
func (usr User) Username() string {
  return usr.username
}

func (usr User) Messages() map[string][]string {
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
  usr.queuedMessages = make(map[string][]string)
}

func (usr *User) AddToQueue(from string, msg string) {
  msgs := usr.queuedMessages[from]

  usr.queuedMessages[from] = append(msgs, msg)
}

func (usr *User) Connect(c net.Conn) {
  usr.isConnected = true
  usr.conn = c
}

func (usr *User) Disconnect() {
  usr.isConnected = false
  usr.conn = nil
}

// ---------------------- value-receiver functions ----------------------------
func (usr User) Connected() bool {
  return usr.isConnected
}

func (usr User) Send(from, mess string) {
  fmt.Fprintf(usr.conn, "Message from %s as follows: %s\r\n", from, mess)
}

func (usr User) GetConn() net.Conn {
  return usr.conn
}

func (usr User) QueueLength() int {
  return len(usr.queuedMessages)
}
