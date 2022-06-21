package studentvue

import "time"

/// This is the time format used by StudentVue
const fixedFormat = "1/2/2006"

/// I need to wrap the Time struct in my own to enable xml encoding and decoding
type Time struct {
	Time time.Time
}

func (t Time) MarshalText() ([]byte, error) {
	text := t.Time.Format(fixedFormat)
	return []byte(text), nil
}

func (t *Time) UnmarshalText(text []byte) error {
	ti, err := time.Parse(fixedFormat, string(text))
	if err == nil {
		t.Time = ti
	}
	return err
}
