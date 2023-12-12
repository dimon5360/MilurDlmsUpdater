package hdlc

import (
	"fmt"
	"log"
)

type IO interface {
	Write([]byte) (int, error)
	Read([]byte) (int, error)
}

var PhysAddrRequest = []byte{0xFA, 0xA5, 0x5F, 0xDC, 0x30, 0x81}

func queryPhysicalAddress(io IO) ([]byte, error) {

	io.Write(PhysAddrRequest)

	resp := make([]byte, 0)
	bytes := make([]byte, 32)

	for {
		n, err := io.Read(bytes)
		if err != nil {
			log.Fatal(err)
			break
		}
		if n == 0 {
			break
		}
		resp = append(resp, bytes[:n]...)
	}

	log.Printf("Got bytes %v\n", resp)

	if [4]byte(resp[:4]) != [4]byte([]byte{0xFA, 0xA5, 0x5F, 0xDC}) {
		return nil, fmt.Errorf("%s", "Physical address doesn't responded")
	}

	addressSize := resp[4]
	adderss := resp[5:5+addressSize]

	return adderss, nil
}