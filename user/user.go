package user

// import (
//   "fmt"
// )

type User struct {
  online bool
  username string
  password string
  queuedMessages []string
}

//func New() User {
//  return User{}
//}

// ----------------------------- getters --------------------------------------
func (usr User) IsOnline() bool {
  return usr.online
}

func (usr User) Username() string {
  return usr.username
}

func (usr User) QueueLength() int {
  return len(usr.queuedMessages)
}

func (usr User) FullQueue() []string {
  return usr.queuedMessages
}

// ---------------------------- functions -------------------------------------

/*
  returns false if login fails, true otherwise
 */
func (usr *User) Login(name string, pass string) bool {
  if usr.username == name && usr.password == pass {
    usr.online = true
    return true
  }

  return false
}

// set user's online status to false
func (usr *User) Logout() {
  usr.online = false
}

// adds message to queue
func (usr *User) AddToQueue(msg string) {
  usr.queuedMessages = append(usr.queuedMessages, msg)
}
