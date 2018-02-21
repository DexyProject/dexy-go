package types

import (
	"testing"
	"time"
)

func TestTimestamp_UnixMilli(t *testing.T) {

	ts := Timestamp{time.Unix(0, 1519213982123000000)}

	if ts.UnixMilli() != 1519213982123 {
		t.Errorf("timestamps do not match")
	}
}
