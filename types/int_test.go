package types

import (
	"testing"
	"strconv"
)

func TestInt_UnmarshalJSON(t *testing.T) {

	val := int64(10)
	i := NewInt(val)

	j, _ := i.MarshalJSON()

	ni := new(Int)

	ni.UnmarshalJSON(j)

	if ni.String() != strconv.Itoa(int(val)) {
		t.Errorf("%s does not equal expected %d", ni.String(), val)
	}
}
