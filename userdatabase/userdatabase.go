/*
  Author: Brendon Rollins, Kyle Brewington, Mason Baird
  Date: Feb 26th 2018

  Description:
    This file defines a basic UserDatabase object. This is essentially a wrapper
    for a map of users with string keys (user names).
*/

package userdatabase

import (
  "user"
  "sync"
)

/*
  UserDatabase object
 */
type UserDatabase struct {
  mutex sync.RWMutex
  users map[string]*user.User
}

/*
  UserDatabase constructor
 */
func NewUserDatabase() *UserDatabase {
  return &UserDatabase{sync.RWMutex{}, map[string]*user.User{}}
}

// ------------------------- value-receiver funcs -----------------------------
/*
  Get user object fro userdatabase
  This function assumes that the user is in the database
 */
func (users UserDatabase) GetUser(name string) *user.User {
  users.mutex.RLock() // lock users for reading
  usr := users.users[name]
  users.mutex.RUnlock() // unlock users for reading

  return usr
}

/*
  Returns true or false if user exsists
 */
func (users UserDatabase) Exists(name string) bool {
  _, exists := users.users[name]

  return exists
}

/*
  Returns number of users added with CRTE
*/
func (users UserDatabase) GetLength() int {
  return len(users.users)
}

// ----------------------- pointer-receiver funcs -----------------------------
/*
  Add a user to the UserDatabase
 */
func (users *UserDatabase) AddUser(n, p string) {
  users.mutex.Lock() // lock for writing

  usr := user.NewUser(n, p)
  users.users[n] = usr

  users.mutex.Unlock() // unlock users for writing
}
