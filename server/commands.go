/*
  Authors: Brendon Rollins, Kyle Brewington, Mason Baird
  Date: Feb 26th 2018

  Description:
    This file contains functions to perform all of the available commands that
    the server performs.
*/

package main

import (
  "fmt"
  "connection"
)

/*
  Create User

  Adds a user object to the user database with the appropriate properties.
  Sends appropriate messages to connection.
 */
func CreateUser(input []string, conn *connection.Connection) {
  name, pass := input[0], input[1]
  superuser := false // superuser flag, first user to be created will be the superuser

  // check if user already exists
  if conn.GetUsers().Exists(name) {
    fmt.Fprintf(conn.GetConn(), "203 User %s already exists.\r\n", name)
    return
  }

  // get userdatabase object (map wrapper)
  usrdatabase := conn.GetUsers()

  // toggle superuser flag if this is the first user created
  if usrdatabase.GetLength() == 0 {
    superuser = true
  }

  // add user to userdatabase
  usrdatabase.AddUser(name, pass)
  fmt.Fprintf(conn.GetConn(), "104 User %s created.\r\n", name)

  // inform connection it is superuser
  if superuser {
    fmt.Fprintf(conn.GetConn(), "105 User %s created as superuser.\r\n", name)
  }
}

/*
  Authenticate User

  Authenticates the connection by tying its session to a user object.
 */
func Authenticate(input []string, conn *connection.Connection) {
  name, pass := input[0], input[1]
  // get user object from connection
  usr := conn.GetUsers().GetUser(name)

  // make sure the user exists
  if !conn.GetUsers().Exists(name) {
    fmt.Fprintf(conn.GetConn(), "200 User %s doesn't exist.\r\n", name)
    return
  }

  // make sure this connection isn't already tied to another user
  if conn.IsAuthorized() {
    name := conn.GetUser().Username()
    fmt.Fprintf(conn.GetConn(), "202 Already connected as %s.\r\n", name)
    return
  }

  // make sure the user is not already authenticated in another connection
  if usr.Connected() {
    fmt.Fprintf(conn.GetConn(), "201 User %s already connected.\r\n", name)
    return
  }

  // make sure credentials match this user
  if !usr.Authenticate(name, pass) {
    fmt.Fprintf(conn.GetConn(), "204 Invalid username or password.\r\n")
    return
  }

  // tie user object to connection
  conn.SetUser(usr)
  usr.Connect(conn.GetConn())
  fmt.Fprintf(conn.GetConn(), "102 Connected as %s\r\n", name)

  // if the user has queued messages (messages send while offline) send them
  // to the connection now
  if usr.QueueLength() > 0 {
    messages := usr.Messages()

    for from, msg := range messages {
      for _, m := range msg {
        fmt.Fprintf(conn.GetConn(), "100 Message from %s as follows: \"%s\"\r\n", from, m)
      }
    }
    usr.ClearMessages()
  }
}

/*
  Send Message

  Send a message from one connection to another
 */
func Send(input []string, conn *connection.Connection) {

  // make sure the sending connection is attached to a user object
  if !conn.IsAuthorized() {
    fmt.Fprintf(conn.GetConn(), "206 Not connected as a user.\r\n")
    return
  }

  to, mess := input[0], input[1:]
  message := reconstructMessage(mess) // convert mess to string

  db := conn.GetUsers()
  if !db.Exists(to) { // make sure the receiving user exists
    fmt.Fprintf(conn.GetConn(), "200 User %s does not exist.\r\n", to)
    return
  }

  to_usr := db.GetUser(to)
  to_conn := to_usr.GetConn()
  from_usr := conn.GetUser().Username()

  // if the receiving user is not connected, add the message to their queue
  if !to_usr.Connected() {
    to_usr.AddToQueue(from_usr, message)
    fmt.Fprintf(conn.GetConn(), "101 Message sent.\r\n")
    return
  }

  fmt.Fprintf(to_conn, "100 Message from %s as follows: \"%s\"\r\n", from_usr, message)
  fmt.Fprintf(conn.GetConn(), "101 Message sent.\r\n")
}

/*
  Quit

  Close the connection and log the user out.

 */
func Quit(input []string, conn *connection.Connection) {
  fmt.Fprintf(conn.GetConn(), "103 Bye.\r\n")
  conn.Close()
}
