/*
  Authors: Brendon Rollins, Kyle Brewington, Mason Baird
  Date: Feb 26th 2018

  Description:
    This file defines a basic connection object.
*/

package connection

import (
  "net"
  "user"
  "userdatabase"
)

/*
  Connection object
 */
type Connection struct {
  conn net.Conn // socket
  user *user.User
  users userdatabase.UserDatabase
  auth bool
}

/*
  Connection constructor
 */
func NewConnection(co net.Conn, usrs userdatabase.UserDatabase) *Connection {
  c := new(Connection)
  c.conn = co
  c.users = usrs
  c.auth = false

  return c
}

// -------------------- pointer-receiver functions ----------------------------
/*
  Disconnect user and close the connection.
 */
func (conn *Connection) Close() {
  conn.user.Disconnect()
  conn.conn.Close()
}

/*
  Set user object for conn and toggle authenticated flag
 */
func (conn *Connection) SetUser(usr *user.User) {
  conn.auth = true
  conn.user = usr
}

// ---------------------- value-receiver functions ----------------------------
/*
  return authorized flag
 */
func (conn Connection) IsAuthorized() bool {
  return conn.auth
}

/*
  return pointer to user object
 */
func (conn Connection) GetUser() *user.User {
  return conn.user
}

/*
  return connection object
 */
func (conn Connection) GetConn() net.Conn {
  return conn.conn
}

/*
  get UserDatabase object
 */
func (conn Connection) GetUsers() userdatabase.UserDatabase {
  return conn.users
}
