package users

import (
	"log"
	"sync"
)

type ConnectedUsers struct {
	Users     []string `json:"users"`
	mx        sync.Mutex
	Listeners []chan []string
}

func (c *ConnectedUsers) Add(user string) {
	c.mx.TryLock()
	defer c.mx.Unlock()
	c.Users = append(c.Users, user)
	for _, l := range c.Listeners {
		l <- c.Users
	}
}

func (c *ConnectedUsers) Remove(user string) {
	c.mx.TryLock()
	defer c.mx.Unlock()
	log.Println("Removing user: " + user)
	for i, u := range c.Users {
		if u == user {
			c.Users = append(c.Users[:i], c.Users[i+1:]...)
			for _, l := range c.Listeners {
				l <- c.Users
			}
			return
		}
	}
}

func (c *ConnectedUsers) Get() []string {
	c.mx.TryLock()
	defer c.mx.Unlock()
	return c.Users
}

func (c *ConnectedUsers) Len() int {
	c.mx.TryLock()
	defer c.mx.Unlock()
	return len(c.Users)
}

func (c *ConnectedUsers) Contains(user string) bool {
	c.mx.TryLock()
	defer c.mx.Unlock()
	for _, u := range c.Users {
		if u == user {
			return true
		}
	}
	return false
}

func (c *ConnectedUsers) AddListener(listener chan []string) {
	c.mx.TryLock()
	defer c.mx.Unlock()
	c.Listeners = append(c.Listeners, listener)
}

func (c *ConnectedUsers) RemoveListener(listener chan []string) {
	c.mx.TryLock()
	defer c.mx.Unlock()
	log.Println("Removing listener")
	for i, l := range c.Listeners {
		if l == listener {
			c.Listeners = append(c.Listeners[:i], c.Listeners[i+1:]...)
			return
		}
	}
}

func NewConnectedUsers() *ConnectedUsers {
	return &ConnectedUsers{
		Users:     []string{},
		Listeners: []chan []string{},
	}
}
