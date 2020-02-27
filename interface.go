package multiprotocol

import (
	"encoding"
	"encoding/json"
)

/*
Multiprotocol is a cross-protocol, cross-platform format for representing
self-describing protocol identifiers.
Learn more here: https://github.com/vacp2p/multiprotocol

Multiprotocol have both a binary and string representation.

    import mp "github.com/vacp2p/go-multiprotocol"

    proto, err := mp.NewMultiprotocol("/vac/waku/2")
    // err non-nil when parsing failed.

*/
type Multiprotocol interface {
	json.Marshaler
	json.Unmarshaler
	encoding.TextMarshaler
	encoding.TextUnmarshaler
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler

	// Equal returns whether two Multiprotocols are exactly equal
	Equal(Multiprotocol) bool

	// Bytes returns the []byte representation of this Multiprotocol
	//
	// This function may expose immutable, internal state. Do not modify.
	Bytes() []byte

	// String returns the string representation of this Multiprotocol
	// (may panic if internal state is corrupted)
	String() string

	// Protocols returns the list of Protocols this Multiprotocol includes
	// will panic if protocol code incorrect (and bytes accessed incorrectly)
	Protocols() []Protocol

	// ValueForProtocol returns the value (if any) following the specified protocol
	//
	// Note: protocols can appear multiple times in a single multiprotocol.
	// Consider using `ForEach` to walk over the addr manually.
	ValueForProtocol(code int) (string, error)
}
