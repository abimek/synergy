package studentvue

import (
	"encoding/xml"
	"strings"
	"time"
)

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

type xmlIntermediary struct {
	Text string `xml:",chardata"`
}

func GetXmlString(text string) (*string, error) {
	sxml := strings.Replace(text, "string", "StudentVueApi", 2)

	x := xmlIntermediary{}
	err := xml.Unmarshal([]byte(sxml), &x)
	if err != nil {
		return nil, err
	}

	return &x.Text, nil
}
