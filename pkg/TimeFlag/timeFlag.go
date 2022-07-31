package TimeFlag

import (
	"fmt"
	"time"
)

const DefaultLayout = "20060102"

// Time
type TimeValue struct {
	Time *time.Time
}

func (v TimeValue) String() string {
	if v.Time != nil {
		return v.Time.Format(DefaultLayout)
	}
	return ``
}

func (v TimeValue) Set(s string) error {

	t, err := time.Parse(DefaultLayout, s)
	if err != nil {
		return fmt.Errorf("Format is not acceptable")
	}
	*v.Time = t
	return nil
}
