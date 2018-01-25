package types

import (
	"testing"
	"encoding/json"
)

var hash = "0x153e75acf6376f9c9c89934f4224710f54a4270a6b552797d49be8685d8880e7"

func TestHash_String(t *testing.T) {
	if hash != NewHash(hash).String() {
		t.Fail()
	}
}

func TestHash_MarshalJSON_UnmarshalJSON(t *testing.T) {
	h := NewHash(hash)

	marshaled, err := json.Marshal(h)
	if err != nil {
		t.Errorf("failed to marshal: %s", err.Error())
	}

	var unmarshal Hash
	err = json.Unmarshal(marshaled, &unmarshal)
	if err != nil {
		t.Errorf("failed to unmarshal: %s", err.Error())
	}

	if unmarshal.String() != h.String() {
		t.Errorf("failed")
	}
}
