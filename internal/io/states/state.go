package state

type State interface {
    Open()
	Idle()
    Exchange()
	Close()

	GetDescription() string
}