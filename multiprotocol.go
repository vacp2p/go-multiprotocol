package multiprotocol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/multiformats/go-multiaddr"
)

// multiprotocol is the data structure representing a Multiprotocol
type multiprotocol struct {
	bytes []byte
}

// NewMultiprotocol parses and validates an input string, returning a *Multiprotocol
func NewMultiprotocol(s string) (p Multiprotocol, err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("Panic in NewMultiaddr on input %q: %s", s, e)
			err = fmt.Errorf("%v", e)
		}
	}()
	b, err := stringToBytes(s)
	if err != nil {
		return nil, err
	}
	return &multiprotocol{bytes: b}, nil
}

// NewMultiprotocolBytes initializes a Multiprotocol from a byte representation.
// It validates it as an input string.
func NewMultiprotocolBytes(b []byte) (a Multiprotocol, err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("Panic in NewMultiaddrBytes on input %q: %s", b, e)
			err = fmt.Errorf("%v", e)
		}
	}()

	if err := validateBytes(b); err != nil {
		return nil, err
	}

	return &multiprotocol{bytes: b}, nil
}

// Equal tests whether two multiprotocols are equal
func (m *multiprotocol) Equal(m2 Multiprotocol) bool {
	return bytes.Equal(m.bytes, m2.Bytes())
}

// Bytes returns the []byte representation of this Multiprotocol
//
// Do not modify the returned buffer, it may be shared.
func (m *multiprotocol) Bytes() []byte {
	return m.bytes
}

// String returns the string representation of a Multiprotocol
func (m *multiprotocol) String() string {
	s, err := bytesToString(m.bytes)
	if err != nil {
		panic(fmt.Errorf("multiprotocol failed to convert back to string. corrupted? %s", err))
	}
	return s
}

func (m *multiprotocol) MarshalBinary() ([]byte, error) {
	return m.Bytes(), nil
}

func (m *multiprotocol) UnmarshalBinary(data []byte) error {
	new, err := NewMultiprotocolBytes(data)
	if err != nil {
		return err
	}
	*m = *(new.(*multiprotocol))
	return nil
}

func (m *multiprotocol) MarshalText() ([]byte, error) {
	return []byte(m.String()), nil
}

func (m *multiprotocol) UnmarshalText(data []byte) error {
	new, err := NewMultiprotocol(string(data))
	if err != nil {
		return err
	}
	*m = *(new.(*multiprotocol))
	return nil
}

func (m *multiprotocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

func (m *multiprotocol) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	new, err := NewMultiprotocol(v)
	*m = *(new.(*multiprotocol))
	return err
}

// Protocols returns the list of protocols this Multiprotocol has.
// will panic in case we access bytes incorrectly.
func (m *multiprotocol) Protocols() []Protocol {
	ps := make([]Protocol, 0, 8)
	b := m.bytes
	for len(b) > 0 {
		code, n, err := multiaddr.ReadVarintCode(b)
		if err != nil {
			panic(err)
		}

		p := ProtocolWithCode(code)
		if p.Code == 0 {
			// this is a panic (and not returning err) because this should've been
			// caught on constructing the Multiaddr
			panic(fmt.Errorf("no protocol with code %d", b[0]))
		}
		ps = append(ps, p)
		b = b[n:]

		n, size, err := sizeForAddr(p, b)
		if err != nil {
			panic(err)
		}

		b = b[n+size:]
	}
	return ps
}

var ErrProtocolNotFound = fmt.Errorf("protocol not found in multiaddr")

func (m *multiprotocol) ValueForProtocol(code int) (value string, err error) {
	err = ErrProtocolNotFound
	ForEach(m, func(c Component) bool {
		if c.Protocol().Code == code {
			value = c.Value()
			err = nil
			return false
		}
		return true
	})
	return
}

