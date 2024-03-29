package types

import (
	"fmt"
	"strconv"
	"time"
)

type Timestamp struct {
	time.Time
}

func NewTimestampFromInt(t Int) Timestamp {
	return Timestamp{time.Unix(t.Int64(), 0)}
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprint(t.Unix())), nil
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}

	t.Time = time.Unix(int64(ts), 0)

	return nil
}
