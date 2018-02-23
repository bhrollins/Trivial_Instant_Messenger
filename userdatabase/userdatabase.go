package userdatabase

import (
  "user"
  "sync"
  // "fmt"
)

type UserDatabase struct {
  mutex sync.RWMutex
  users map[string]*user.User
}

// constructor
func NewUserDatabase() *UserDatabase {
  return &UserDatabase{sync.RWMutex{}, map[string]*user.User{}}
}

// ------------------------- value-receiver funcs -----------------------------
func (users UserDatabase) GetUser(name string) *user.User {
  users.mutex.RLock()
  usr := users.users[name]
  users.mutex.RUnlock()

  return usr
}

func (users UserDatabase) Exists(name string) bool {
  _, exists := users.users[name]

  return exists
}

func (users UserDatabase) GetLength() int {
  return len(users.users)
}

// ----------------------- pointer-receiver funcs -----------------------------
func (users *UserDatabase) AddUser(n, p string) {
  users.mutex.Lock() // lock for writing

  usr := user.NewUser(n, p)
  users.users[n] = usr

  users.mutex.Unlock()
}
