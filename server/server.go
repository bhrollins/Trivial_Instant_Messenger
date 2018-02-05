package main

import (
  "net"
  "fmt"
  // "time"
  "bufio"
  "strings"
  "user"
  "usermanager"
)

// CRTE -- create user account (provide desired username and password)
func createUser(original string, input []string) string {

  if len(input) != 2 {
    return fmt.Sprintf("Usage: CRTE [username] [password]")
  }

  // check if user exists
  if usermanager.Exists(input[0]) {
    return fmt.Sprintf("User already exists")
  }

  // create user
  usr := user.New(false, input[0], input[1])
  usermanager.AddUser(usr)

  return fmt.Sprintf("105 User %s created", input[0])
}

// AUTH -- authenticate a user for a particular account (provide username and password)
func authenticateUser(original string, input []string) string {

  if len(input) != 2 {
    return fmt.Sprintf("Usage: AUTH [username] [password]")
  }

  // check if user exists
  if usermanager.Exists(input[0]) {
    return fmt.Sprintf("User already exists")
  }

  // check if authenticated
  usr := usermanager.GetUser(input[0])

  if user.Login(usr) {
    return fmt.Sprintf("102 Connected as %s", input[0])
  }

  return "Authentication failed"
}

// SEND -- send a message to another user. (the receiving user will be the first word, then the
//         rest of the line is interpreted as the message to be sent) the sending user must
//         have been authenticated and the receiving user must exist. the message will be
//         sent to the receiving user's connection. if the receiving user is not connected the
//         message will be stored to be sent the next time the receiving user logs in
func sendMessage(original string, input []string) string {

  if len(input) < 2 {
    return fmt.Sprintf("Usage: SEND [username] <message>")
  }

  // check if <sending> user is authenticated

  // check if receiving user exists

  // check if receiving user is online

  return "msg"
}

// QUIT -- will disconnect the user from the system cleanly. the connection should also be
//         cleanly terminated if the connection is closed inn some otther way
func quit(original string, input []string) string {
  return "quit"
}

func handleConnection(conn net.Conn, client_id int) {

  fmt.Println("New Connection")

  defer func() {
    fmt.Println("Closing Connection")
    conn.Close()
  }()

  bufReader := bufio.NewReader(conn)

  for {

    input, err := bufReader.ReadString('\n')

    if err != nil {
      fmt.Printf("error: %s\n", err)
      return
    }

    command := strings.Split(input, " ")

    switch strings.TrimSpace(command[0]) {
    case "CRTE":
      output := createUser(input, command[1:])
      fmt.Fprintf(conn, "%s\n", output)
    case "AUTH":
      output := authenticateUser(input, command[1:])
      fmt.Fprintf(conn, "%s\n", output)
    case "SEND":
      output := sendMessage(input, command[1:])
      fmt.Fprintf(conn, "%s\n", output)
    case "QUIT":
      output := quit(input, command[1:])
      fmt.Fprintf(conn, "%s\n", output)
    default:
      fmt.Fprintf(conn, "No such command \"%s\"\n", strings.TrimSpace(input))
    }

  }
}

func main() {

  ln, err := net.Listen("tcp", ":8080")
  clients := 0
  usrMgr := new(usermanager.UserManager)

  if err != nil {
    fmt.Println(err)
    return
  }

  defer func() {
    ln.Close()
    fmt.Println("Listener closed")
  }()

  for {
    conn, err := ln.Accept()

    if err != nil {
      fmt.Println(err)
      return
    }

    go handleConnection(conn, clients)

    clients += 1

  }

}
