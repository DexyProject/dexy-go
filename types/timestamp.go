package types

import (
	"time"
	"fmt"
	"strconv"
)

type Timestamp struct {
	time.Time
}

func UnixMilli(millie int64) Timestamp {
	return Timestamp{time.Unix(0, millie*int64(time.Millisecond))}
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprint(t.UnixMilli())), nil
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}

	t.Time = time.Unix(0, int64(ts)*int64(time.Millisecond))

	return nil
}

func (t Timestamp) UnixMilli() int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
