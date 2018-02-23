package main

import (
  "fmt"
)

func cleanInput(c rune) bool {
  // we only want to accept input in the ascii range of unicode
  return c < 32 || c > 126
}

func removeSpaces(input []string) []string {
  nospace := make([]string, 0)

  for _, val := range input {
    if val != "" {
      nospace = append(nospace, val)
    }
  }

  return nospace
}

func checkInput(command string, input []string) string {
    input = removeSpaces(input)
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
      // TODO: implement check?
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
