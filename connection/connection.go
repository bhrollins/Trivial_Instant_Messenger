package connection

import (
  "net"
  "user"
  "userdatabase"
)

type Connection struct {
  conn net.Conn // socket
  user *user.User
  users userdatabase.UserDatabase
  auth bool
}

// constructor
func NewConnection(co net.Conn, usrs userdatabase.UserDatabase) *Connection {
  c := new(Connection)
  c.conn = co
  c.users = usrs
  c.auth = false

  return c
}

// -------------------- pointer-receiver functions ----------------------------
func (conn *Connection) Close() {
  // TODO: make sure the "gracefully closed" requirement is satisfied
  conn.user.Disconnect()
  conn.conn.Close()
}

func (conn *Connection) SetUser(usr *user.User) {
  conn.auth = true
  conn.user = usr
}

// ---------------------- value-receiver functions ----------------------------
func (conn Connection) IsAuthorized() bool {
  return conn.auth
}

func (conn Connection) GetUser() *user.User {
  return conn.user
}

func (conn Connection) GetConn() net.Conn {
  return conn.conn
}

func (conn Connection) GetUsers() userdatabase.UserDatabase {
  return conn.users
}

// func (conn connection) Send(message string) {
//   conn.user.Send(conn.user.Username(), message)
// }
