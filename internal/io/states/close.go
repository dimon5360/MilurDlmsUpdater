package state

import (
	"log"
)

type ClosedState struct {
}

func (s *ClosedState) Open() {
	log.Fatal("Can't open from Close state")
}

func (s *ClosedState) Idle() {
	log.Fatal("Can't idle from Close state")
}

func (s *ClosedState) Exchange() {
	log.Fatal("Can't exchange from Close state")
}

func (s *ClosedState) Close() {
	log.Println("ClosedState::Close()")
}

func (s *ClosedState) GetDescription() string {
	return "ClosedState"
}