package types

import (
	"fmt"
	"strconv"
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Timestamp struct {
	time.Time
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

func (t Timestamp) GetBSON() (interface{}, error) {
	if t.Time.IsZero() {
		return nil, nil
	}

	return t.Time, nil
}

func (t *Timestamp) SetBSON(raw bson.Raw) error {
	var tm time.Time

	err := raw.Unmarshal(&tm)
	if err != nil {
		return err
	}

	t.Time = tm

	return nil
}
