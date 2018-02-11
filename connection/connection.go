package connection

import (
  "net"
  "user"
  "userdatabase"
)

type Connection struct {
  conn net.Conn // socket
  user user.User
  users userdatabase.UserDatabase
}

// constructor
func NewConnection(co net.Conn, usrs userdatabase.UserDatabase) *Connection {
  c := new(Connection)
  c.conn = co
  c.users = usrs

  return c
}

// -------------------- pointer-receiver functions ----------------------------
func (conn *Connection) Close() {
  // close this gracefully somehow
  // TODO... implement this.
  conn.conn = nil
  // conn.user = nil
  // conn.users = nil
}

func (conn *Connection) SetUser(usr user.User) {
  conn.user = usr
}

// ---------------------- value-receiver functions ----------------------------
func (conn Connection) GetUser() user.User {
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
