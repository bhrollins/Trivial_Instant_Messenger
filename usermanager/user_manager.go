package usermanager

import (
//  "fmt"
  "user"
)

type UserManager struct {
  users map[string]User // maybe??
}

// add a user to the user manager (database)
func (mgr *UserManager) AddUser(usr User) {
  users[user.Username(usr)] = usr
}

// check if username is already taken
func (mgr UserManager) Exists(usr User) bool {
  _, exists := mgr[user.Username(usr)]

  return exists
}
