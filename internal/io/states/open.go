package state

import (
	"log"
)

type OpenState struct {
	Description string
}

func (s *OpenState) Open() {
	log.Println("OpenState::Open()")
}

func (s *OpenState) Idle() {
	log.Fatal("Can't idle from Open state")
}

func (s *OpenState) Exchange() {
	log.Fatal("Can't exchange from Open state")
}

func (s *OpenState) Close() {
	log.Fatal("Can't close from Open state")
}

func (s *OpenState) GetDescription() string {
	return "OpenState"
}