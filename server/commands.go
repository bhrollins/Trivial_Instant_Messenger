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
  superuser := false

  if conn.GetUsers().Exists(name) {
    fmt.Fprintf(conn.GetConn(), "203 User %s already exists.\r\n", name)
    return
  }

  usrdatabase := conn.GetUsers()

  if usrdatabase.GetLength() == 0 {
    superuser = true
  }

  usrdatabase.AddUser(name, pass)
  fmt.Fprintf(conn.GetConn(), "104 User %s created.\r\n", name)

  if superuser {
    fmt.Fprintf(conn.GetConn(), "105 User %s created as superuser.\r\n", name)
  }
}

/*
  Authenticate User
 */
func Authenticate(input []string, conn *connection.Connection) {
  name, pass := input[0], input[1]
  usr := conn.GetUsers().GetUser(name)

  if !conn.GetUsers().Exists(name) {
    fmt.Fprintf(conn.GetConn(), "200 User %s doesn't exist.\r\n", name)
    return
  }

  if conn.IsAuthorized() {
    name := conn.GetUser().Username()
    fmt.Fprintf(conn.GetConn(), "202 Already connected as %s.\r\n", name)
    return
  }

  if usr.Connected() {
    fmt.Fprintf(conn.GetConn(), "201 User %s already connected.\r\n", name)
    return
  }

  if !usr.Authenticate(name, pass) {
    fmt.Fprintf(conn.GetConn(), "204 Invalid username or password.\r\n")
    return
  }

  conn.SetUser(usr)
  usr.Connect(conn.GetConn())
  fmt.Fprintf(conn.GetConn(), "102 Connected as %s\r\n", name)

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
 */
func Send(input []string, conn *connection.Connection) {

  if !conn.IsAuthorized() {
    fmt.Fprintf(conn.GetConn(), "206 Not connected as a user.\r\n")
    return
  }

  to, mess := input[0], input[1:]
  message := reconstructMessage(mess)

  db := conn.GetUsers()
  if !db.Exists(to) {
    fmt.Fprintf(conn.GetConn(), "200 User %s does not exist.\r\n", to)
    return
  }

  to_usr := db.GetUser(to)
  to_conn := to_usr.GetConn()
  from_usr := conn.GetUser().Username()

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
 */
func Quit(input []string, conn *connection.Connection) {
  fmt.Fprintf(conn.GetConn(), "103 Bye.\r\n")
  conn.Close()
}
