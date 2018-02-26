/*
  Authors: Brendon Rollins, Kyle Brewington, Mason Baird
  Date: Feb 26th 2018

  Description:
    This file provides helpers functions for the server.
*/

package main

import (
  "fmt"
)

/*
  Cleans input so only ascii range of unicode is accepted
*/
func cleanInput(c rune) bool {
  return c < 32 || c > 126
}

/*
  remove white space from input
*/
func removeSpaces(input []string) []string {
  nospace := make([]string, 0)

  for _, val := range input {
    if val != "" {
      nospace = append(nospace, val)
    }
  }

  return nospace
}

/*
  evaluates user input and runs proper case
*/
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
      return ""
    }
    return ""
}

/*
  Reconstructs message and removes implicit space at end of messages
*/
func reconstructMessage(mess []string) string {
  str_out := ""
  for _, each := range mess {
    str_out += each + " "
  }

  str_out = str_out[:len(str_out)-1]

  return str_out
}
