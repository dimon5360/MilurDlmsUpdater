package state

import (
	"log"
)

type ExchangeState struct {
	Description string
}

func (s *ExchangeState) Open() {
	log.Fatal("Can't open from Exchange state")
}

func (s *ExchangeState) Idle() {
	log.Fatal("Can't idle from Exchange state")
}

func (s *ExchangeState) Exchange() {
	log.Println("ExchangeState::Exchange()")
}

func (s *ExchangeState) Close() {
	log.Fatal("Can't close from Close state")
}

func (s *ExchangeState) GetDescription() string {
	return "ExchangeState"
}