package usermanager

import (
//  "fmt"
  . "user"
)

type UserManager struct {
  users map[string]User // maybe??
}

// add a user to the user manager (database)
func (mgr *UserManager) AddUser(usr User) {
  mgr.users[usr.Username()] = usr
}

// check if username is already taken
func (mgr UserManager) Exists(usr User) bool {
  _, exists := mgr.users[usr.Username()]

  return exists
}
