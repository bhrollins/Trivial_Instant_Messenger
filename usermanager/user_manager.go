package usermanager

import (
    // "fmt"
  . "user"
)

type UserManager struct {
  users map[string]User
}

// add a user to the user manager (database)
func (mgr *UserManager) AddUser(usr User) {
  if mgr.users == nil {
    mgr.users = make(map[string]User)
  }
  mgr.users[usr.Username()] = usr
}

// check if username is already taken
func (mgr UserManager) Exists(usr User) bool {
  _, exists := mgr.users[usr.Username()]

  return exists
}

func (mgr UserManager) GetUser(username string) User {
  return mgr.users[username]
}
