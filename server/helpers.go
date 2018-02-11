package main

import (
  "fmt"
)

func cleanInput(c rune) bool {
  // we only want to accept input in the ascii range of unicode
  return c < 32 || c > 126
}

func checkInput(command string, input []string) string {
    switch command {
    case "CRTE":
      fallthrough
    case "AUTH":
      if len(input) != 2 {
        return fmt.Sprintf("Usage: %s [username] [password]", command)
      }
    case "SEND":
      if len(input) < 2 {
        return fmt.Sprintf("Usage: SEND [username] <message>")
      }
    case "QUIT":
      // todo implement check
      return ""
    }

    return ""
}

func reconstructMessage(mess []string) string {
  str_out := ""
  for _, each := range mess {
    str_out += each + " "
  }

  return str_out
}
