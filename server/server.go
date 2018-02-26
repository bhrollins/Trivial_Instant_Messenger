/*
  Authors: Brendon Rollins, Kyle Brewington, Mason Baird
  Date: Feb 26th 2018

  Description:
    This is the entry point for this project. Here we start a socket listening
    on port 7778 and spin off a thread for each connection made.
*/

package main

import (
  "fmt"
  "net"
  "bufio"
  "strings"
  "userdatabase"
  "connection"
)

// basic server struct
type server struct {
  users userdatabase.UserDatabase
}

// global server object
var srvr = new(server)

// ------------------------------ handle connection ---------------------------
/*
  handleConnection

  thread target function
    accepts input from connection and executes appropriate commands
 */
func handleConnection(conn net.Conn) {
  // get ip of connection
  ip := conn.RemoteAddr().(*net.TCPAddr).IP

  c := connection.NewConnection(conn, srvr.users)
  fmt.Printf("New Connection from: %s\n", ip)

  // close connection when this function exits
  defer func() {
    fmt.Printf("Closing connection for: %s\n", ip)
    conn.Close()
  }()

  bufReader := bufio.NewReader(conn)

  // infinite loop while this connection is live
  for {
    input, err := bufReader.ReadString('\n')

    if err != nil {
      fmt.Printf("Error while reading string: %s\n", err)
      return
    }

    // remove any non-ascii characters:
    input = strings.TrimFunc(input, cleanInput)
    // trim leading and tailing whitespace:
    input = strings.Trim(input, " ")
    fullCommand := strings.Split(input, " ")
    command, args := fullCommand[0], fullCommand[1:]

    str_err := checkInput(command, args)

    if str_err != "" {
      fmt.Fprintf(conn, "%s\r\n", str_err)
      continue
    }

    // Depending on the command, do said function
    switch command {
    case "CRTE":
      CreateUser(args, c)
    case "AUTH":
      Authenticate(args, c)
    case "SEND":
      Send(args, c)
    case "QUIT":
      Quit(args, c)
    default:
      fmt.Fprintf(conn, "205 No such command \"%s\".\r\n", input)
    }

  }

}

// ---------------------------------- main ------------------------------------
func main() {
  // instantiate server properties
  db := userdatabase.NewUserDatabase()
  srvr.users = *db

  // create socket listening on port 7778
  ln, err := net.Listen("tcp", ":7778")

  // If error, print error and exit
  if err != nil {
    fmt.Println(err)
    return
  }

  // called when main exits:
  defer func() {
    ln.Close()
    fmt.Println("Listener closed")
  }()

  // infinite loop
  for {
    conn, err := ln.Accept()

    if err != nil {
      fmt.Println(err)
      continue
    }

    go handleConnection(conn)

  }

}
