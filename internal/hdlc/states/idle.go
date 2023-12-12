package state

import (
	"log"
)

type IdleState struct {
	Description string
}

func (s *IdleState) Open() {
	log.Fatal("Can't open from Idle state")
}

func (s *IdleState) Idle() {
	log.Println("IdleState::Idle()")
}

func (s *IdleState) Exchange() {
	log.Fatal("Can't exchange from Idle state")
}

func (s *IdleState) Close() {
	log.Fatal("Can't close from Close state")
}

func (s *IdleState) GetDescription() string {
	return "IdleState"
}