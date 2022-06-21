package gradebook

import (
	"fmt"
	"strconv"
	"strings"
)

type GradeType string

const (
	MajorSummative  GradeType = "Major Summative"
	MinorSummative  GradeType = "Minor Summative"
	FormalFormative GradeType = "Formal Formative"
)

type PointType string

const (
	NormalPoints  PointType = "normal"
	PendingPoints PointType = "pending"
)

type ScoreType string

const (
	NormalScore          ScoreType = "normal"
	NotGradedScore       ScoreType = "not graded"
	CharachterGradeScore ScoreType = "charachter grade"
)

type Points struct {
	Type   PointType
	Points float64
	Total  float64
}

type Score struct {
	Type        ScoreType
	Points      float64
	Total       float64
	LetterGrade string
}

// Marshal for Score struct
func (f Score) MarshalText() ([]byte, error) {
	var text string
	if f.Type == NormalScore {
		text = fmt.Sprintf("%f", f.Points) + " out of " + fmt.Sprintf("%f", f.Total) + " (" + f.LetterGrade + ")"
	}
	if f.Type == NotGradedScore {
		text = "Not Graded"
	}
	if f.Type == CharachterGradeScore {
		text = f.LetterGrade
	}
	return []byte(text), nil
}

// Unmarshal for Score struct
func (f *Score) UnmarshalText(text []byte) error {
	t := string(text)
	if t == "Not Graded" {
		f.Type = NotGradedScore
		f.Points = -1.0
		f.Total = -1.0
		f.LetterGrade = "N/A"
		return nil
	}
	if len(strings.Split(t, " ")) == 1 {
		f.Type = CharachterGradeScore
		f.Points = -1.0
		f.Total = -1.0
		f.LetterGrade = t
		return nil
	}
	f.Type = NormalScore
	components := strings.Split(t, " ")
	p, err := strconv.ParseFloat(components[0], 32)
	if err != nil {
		return err
	}
	f.Points = p
	s, err := strconv.ParseFloat(components[3], 32)
	if err != nil {
		return err
	}
	f.Total = s
	f.LetterGrade = strings.ReplaceAll(strings.ReplaceAll(components[4], "(", ""), ")", "")
	return nil
}

// Marshal for Points
func (f Points) MarshalText() ([]byte, error) {
	text := fmt.Sprintf("%f", f.Points) + " / " + fmt.Sprintf("%f", f.Total)
	return []byte(text), nil
}

// Unmarshal for Points
func (f *Points) UnmarshalText(text []byte) error {
	t := string(text)
	components := strings.Split(t, " / ")
	if len(components) == 1 {
		components = strings.Split(t, " ")
		v, err := strconv.ParseFloat(components[0], 32)
		if err != nil {
			return err
		}
		f.Type = PendingPoints
		f.Points = -1.0
		f.Total = v
		return nil
	}
	f.Type = NormalPoints
	p, err := strconv.ParseFloat(components[0], 32)
	if err != nil {
		return err
	}
	f.Points = p
	p, err = strconv.ParseFloat(components[1], 32)
	if err != nil {
		return err
	}
	f.Total = p
	return nil
}

func (f GradeType) MarshalText() ([]byte, error) {
	return []byte(f), nil
}

func (f *GradeType) UnmarshalText(text []byte) error {
	*f = GradeType(string(text))
	return nil
}
