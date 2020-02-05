package cryptolib

import (
	"testing"
)

// TestEncode comment
func TestEncode(t *testing.T) {
	fakeScript := "fake script"
	fakeScriptBytes := []byte(fakeScript)

	s := EncodeBIP276(prefixScript, networkMainnet, currentVersion, fakeScriptBytes)

	expected := "bitcoin-script:010166616b65207363726970746f0cd86a"

	if s != expected {
		t.Errorf("Expected: %q, got: %q", expected, s)
	}

}

// TestDecode comment
func TestDecode(t *testing.T) {
	testScript := "bitcoin-script:010166616b65207363726970746f0cd86a"

	prefix, network, version, data, err := DecodeBIP276(testScript)
	if err != nil {
		t.Errorf("Error decoding bip276 string: %+v", err)
	} else {
		t.Logf("prefix: %q, network: %d, version: %d, data,: %s", prefix, network, version, data)

	}

}