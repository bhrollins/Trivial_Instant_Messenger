package main

import (
  "fmt"
  "connection"
)

/*
  Create User
 */
func CreateUser(input []string, conn *connection.Connection) {
  name, pass := input[0], input[1]

  if conn.GetUsers().Exists(name) {
    fmt.Fprintf(conn.GetConn(), "User already exists\r\n")
  }

  usrdatabase := conn.GetUsers()
  usrdatabase.AddUser(name, pass)
  // TODO change the below
  fmt.Fprintf(conn.GetConn(), "105 User %s created\r\n", name)
}

/*
  Authenticate User
 */
func Authenticate(input []string, conn *connection.Connection) {
  name, pass := input[0], input[1]
  usr := conn.GetUsers().GetUser(name)

  if !conn.GetUsers().Exists(name){
    fmt.Fprintf(conn.GetConn(), "User does not exist\r\n")
    return
  }

  if !usr.Authenticate(name, pass) {
    fmt.Fprintf(conn.GetConn(), "Invalid password\r\n")
    return
  }

  conn.SetUser(*usr)
  usr.Connect(conn.GetConn())
  fmt.Fprintf(conn.GetConn(), "102 Connected as %s\r\n", name)
}

/*
  Send Message
 */
func Send(input []string, conn *connection.Connection) {
  to, mess := input[0], input[1:]
  message := reconstructMessage(mess)

  to_conn := conn.GetUsers().GetUser(to).GetConn()
  from_usr := conn.GetUser().Username()

  fmt.Fprintf(to_conn, "Message from %s as follows: %s\r\n", from_usr, message)
}

/*
  Quit
 */
func Quit(input []string, conn *connection.Connection) {
  fmt.Fprintf(conn.GetConn(), "Closing connection\r\n")
}
