package hdlc

type IO interface {
	Write([]byte) (int, error)
	Read([]byte) (int, error)
}