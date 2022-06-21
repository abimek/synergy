package gradebook

import (
	"strconv"
	"strings"
)

type Score struct {
	Type        GradeType
	Points      int
	Total       int
	LetterGrade string
}

func (f Score) MarshalText() ([]byte, error) {
	var text string
	if f.Type == Normal {
		text = strconv.Itoa(f.Points) + " out of " + strconv.Itoa(f.Total) + " (" + f.LetterGrade + ")"
	}
	if f.Type == NotGraded {
		text = "Not Graded"
	}
	if f.Type == CharachterGrade {
		text = f.LetterGrade
	}
	return []byte(text), nil
}

func (f *Score) UnmarshalText(text []byte) error {
	t := string(text)
	if t == "Not Graded" {
		f.Type = NotGraded
		f.Points = -1
		f.Total = -1
		f.LetterGrade = "N/A"
		return nil
	}
	if len(strings.Split(t, " ")) == 1 {
		f.Type = CharachterGrade
		f.Points = -1
		f.Total = -1
		f.LetterGrade = t
		return nil
	}
	f.Type = Normal
	components := strings.Split(t, " ")
	p, err := strconv.ParseFloat(components[0], 32)
	if err != nil {
		return err
	}
	f.Points = int(p)
	s, err := strconv.ParseFloat(components[3], 32)
	if err != nil {
		return err
	}
	f.Total = int(s)
	f.LetterGrade = strings.ReplaceAll(strings.ReplaceAll(components[4], "(", ""), ")", "")
	return nil
}
