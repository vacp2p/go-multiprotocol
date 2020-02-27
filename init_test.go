package multiprotocol

import (
	"testing"
)

func TestInit(t *testing.T) {

	err := Init("testdata/multiprotocol.csv")
	if err != nil {
		t.Errorf("unexpected failure: %s", err.Error())
	}

	if len(Protocols) != 4 {
		t.Errorf("unexpected amount of protocols parsed. expected: 4 got: %d", len(Protocols))
	}
}
