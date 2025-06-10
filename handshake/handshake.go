package handshake

import (
	"fmt"
	"io"
)

// a handshake is a special message that a peer uses to id itself
type Handshake struct {
	Pstr     string
	Infohash [20]byte
	PeerId   [20]byte
}

// New creates a new handshake with the std pstr
func New(infoHash, peerId [20]byte) *Handshake {
	return &Handshake{
		Pstr:     "BitTorrent Protocol",
		Infohash: infoHash,
		PeerId:   peerId,
	}
}

// Serialize serializes a handshake to a buffer
func (h *Handshake) Serialize() []byte {
	buf := make([]byte, len(h.Pstr)+49)
	buf[0] = byte(len(h.Pstr))
	curr := 1
	curr += copy(buf[curr:], h.Pstr)
	curr += copy(buf[curr:], make([]byte, 8))
	curr += copy(buf[curr:], h.Infohash[:])
	curr += copy(buf[curr:], h.PeerId[:])
	return buf
}

// Read parses Handshake from a stream
func Read(r io.Reader) (*Handshake, error) {
	bufLen := make([]byte, 1)
	_, err := io.ReadFull(r, bufLen)
	if err != nil {
		return nil, err
	}
	pstrLen := int(bufLen[0])

	if pstrLen == 0 {
		err := fmt.Errorf("pstrlen cannot be 0")
		return nil, err
	}

	handshakeBuf := make([]byte, 48+pstrLen)
	_, err = io.ReadFull(r, handshakeBuf)
	if err != nil {
		return nil, err
	}

	var infoHash, peerId [20]byte

	copy(infoHash[:], handshakeBuf[pstrLen+8:pstrLen+8+20])
	copy(peerId[:], handshakeBuf[pstrLen+8+20:])

	h := &Handshake{
		Pstr:     string(handshakeBuf[0:pstrLen]),
		Infohash: infoHash,
		PeerId:   peerId,
	}

	return h, nil
}
