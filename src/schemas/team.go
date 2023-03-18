package schemas

import (
	"errors"

	"github.com/google/uuid"
)

type Player struct {
	ID         int
	Name       string `json:"Player"`
	ExternalID string `json:"-"`
}

type Team struct {
	ID      int
	Name    string
	Players []Player
	lastID  int
}

func NewTeam() *Team {
	return &Team{
		ID:      1,
		Name:    "Brazil",
		Players: []Player{},
		lastID:  0,
	}
}

func (t *Team) AddPlayer(name string) error {
	if t.lastID >= 11 {
		return errors.New("team not accepting more players")
	}

	newID := t.lastID + 1

	player := Player{
		ID:         newID,
		Name:       name,
		ExternalID: uuid.NewString(),
	}
	t.lastID = newID
	t.Players = append(t.Players, player)

	return nil
}
