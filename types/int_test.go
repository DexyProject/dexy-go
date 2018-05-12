package types

import (
	"testing"
	"strconv"
	"fmt"
)

func TestInt_UnmarshalJSON(t *testing.T) {
	val := 10
	i := new(Int)

	bytes := []byte(fmt.Sprintf("\"%d\"", val))

	i.UnmarshalJSON(bytes)

	if i.String() != strconv.Itoa(val) {
		t.Errorf("%s does not equal expected %d", i.String(), val)
	}
}

func TestInt_MarshalJSON(t *testing.T) {
	val := int64(10)
	i := NewInt(val)

	str, err := i.MarshalJSON()
	if err != nil {
		t.Error(err)
	}

	expected := fmt.Sprintf("\"%d\"", val)

	if string(str) != expected {
		t.Errorf("%s does not equal to expected %s", string(str), expected)
	}
}
