package domain

import "fmt"

type Player struct {
	Name     string `json:"name"`
	Trophies int    `json:"trophies"`
}

func (p *Player) String() string {
	return fmt.Sprintf("Player %s wants to join.\nTotal trophies: %d", p.Name, p.Trophies)
}
