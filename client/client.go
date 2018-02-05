package main

import (
  "net"
  "fmt"
  "bufio"
  "os"
)

func main() {

  conn, err := net.Dial("tcp", "localhost:8080")

  if err != nil {
    fmt.Println(err)
  }

  cl_reader := bufio.NewReader(os.Stdin)
  net_reader := bufio.NewReader(conn)

  for {
    fmt.Print("==> ")
    text, err := cl_reader.ReadString('\n')

    if err != nil {
      fmt.Printf("error: %s\n", err)
      break
    }

    fmt.Fprintf(conn, "%s", text)

    recv, err := net_reader.ReadString('\n')

    if err != nil {
      fmt.Printf("error: %s\n", err)
      break
    }

    fmt.Printf("received: %s\n", recv)

  }

}
