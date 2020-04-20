package script

import (
	"encoding/hex"
	"github.com/libsv/libsv/script"
	"testing"
)

func TestIsPublicKeyHashOut(t *testing.T) {
	b, _ := hex.DecodeString("76a91403ececf2d12a7f614aef4c82ecf13c303bd9975d88ac")
	scriptPub := script.NewScriptFromBytes(b)

	ret := scriptPub.IsPublicKeyHashOut()

	t.Log(ret)
}

func TestIsPublicKeyOut(t *testing.T) {
	b, _ := hex.DecodeString("2102f0d97c290e79bf2a8660c406aa56b6f189ff79f2245cc5aff82808b58131b4d5ac")
	scriptPub := script.NewScriptFromBytes(b)

	ret := scriptPub.IsPublicKeyOut()

	t.Log(ret)
}

func TestIsScriptHashOut(t *testing.T) {
	b, _ := hex.DecodeString("a9149de5aeaff9c48431ba4dd6e8af73d51f38e451cb87")
	scriptPub := script.NewScriptFromBytes(b)

	ret := scriptPub.IsScriptHashOut()

	t.Log(ret)
}

func TestIsMultisigOut(t *testing.T) {
	b, _ := hex.DecodeString("5201110122013353ae")
	scriptPub := script.NewScriptFromBytes(b)

	ret := scriptPub.IsMultisigOut()

	t.Log(ret)
}

func TestGetPublicKeyHash(t *testing.T) {
	b, _ := hex.DecodeString("76a91404d03f746652cfcb6cb55119ab473a045137d26588ac")
	s := script.NewScriptFromBytes(b)

	pkh, err := s.GetPublicKeyHash()
	if err != nil {
		t.Error(err)
	}

	expected := "04d03f746652cfcb6cb55119ab473a045137d265"

	if hex.EncodeToString(pkh) != expected {
		t.Fail()
	}
	// t.Logf("%x\n", pkh)
}

func TestGetPublicKeyHashAsString(t *testing.T) {
	s := script.NewScriptFromString("76a91404d03f746652cfcb6cb55119ab473a045137d26588ac")

	pkh, err := s.GetPublicKeyHash()
	if err != nil {
		t.Error(err)
	}

	expected := "04d03f746652cfcb6cb55119ab473a045137d265"

	if hex.EncodeToString(pkh) != expected {
		t.Fail()
	}
	// t.Logf("%x\n", pkh)
}

func TestGetPublicKeyHashEmptyScript(t *testing.T) {
	s := script.NewScript()

	_, err := s.GetPublicKeyHash()
	if err == nil {
		t.Error("Expected 'Script is empty'")
	}
}