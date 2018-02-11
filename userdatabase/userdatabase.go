package userdatabase

import (
  "user"
  // "fmt"
)

type UserDatabase struct {
  // TODO: make this safe for concurrent reads/writes
  users map[string]*user.User
}

// constructor
func NewUserDatabase() *UserDatabase {
  return &UserDatabase{map[string]*user.User{}}
}

// ------------------------- value-receiver funcs -----------------------------
func (users UserDatabase) GetUser(name string) *user.User {
  return users.users[name]
}

func (users UserDatabase) Exists(name string) bool {
  _, exists := users.users[name]

  return exists
}

// ----------------------- pointer-receiver funcs -----------------------------
func (users *UserDatabase) AddUser(n, p string) {
  usr := user.NewUser(n, p)

  users.users[n] = usr
}
