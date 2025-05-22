package server

import (
	"math/rand"
	"sync"
)

type UserAgents struct {
	mu     sync.Mutex
	agents []string
}

func NewUserAgents(agents []string) *UserAgents {
	return &UserAgents{
		agents: agents,
		mu:     sync.Mutex{},
	}
}

func (u *UserAgents) Random() string {
	u.mu.Lock()
	defer u.mu.Unlock()
	return u.agents[rand.Intn(len(u.agents))]
}
