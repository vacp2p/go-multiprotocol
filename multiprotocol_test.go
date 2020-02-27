package multiprotocol_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/vacp2p/go-multiprotocol"
)

func TestMain(m *testing.M) {
	multiprotocol.Init("testdata/multiprotocol.csv")
	os.Exit(m.Run())
}

func TestMultiprotocol_ValueForProtocol(t *testing.T) {
	expected := "2"

	mp, err := multiprotocol.NewMultiprotocol(fmt.Sprintf("/vac/waku/%s", expected))
	if err != nil {
		t.Errorf("failed to create multiprotocol: %s", err.Error())
	}

	protocol := multiprotocol.ProtocolWithName("waku")

	val, err := mp.ValueForProtocol(protocol.Code)
	if err != nil {
		t.Errorf("failed to get value %s", err.Error())
	}

	if val != expected {
		t.Errorf("did not match expected: %s actual: %s", expected, val)
	}
}

func TestMultiprotocol_StringReturnsExpected(t *testing.T) {
	str := "/vac/waku/2/store/2/relay/2"

	mp, _ := multiprotocol.NewMultiprotocol(str)
	m, _ := multiprotocol.NewMultiprotocolBytes(mp.Bytes())

	if str != m.String() {
		t.Errorf("strings did not match expected: %s actual %s", str, m.String())
	}
}
