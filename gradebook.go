package studentvue

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type GradeBook struct {
	Type                       string         `xml:"Type,attr"`
	ErrorMessage               string         `xml:"ErrorMessage,attr"`
	HideStandardGraphInd       bool           `xml:"HideStandardGraphInd,attr"`
	HideMarksColumnElementary  bool           `xml:"HideMarksColumnElementary,attr"`
	HidePointsColumnElementary bool           `xml:"HidePointsColumnElementary,attr"`
	HidePercentSecondary       bool           `xml:"HidePercentSecondary,attr"`
	DisplayStandardsData       bool           `xml:"DisplayStandardsData,attr"`
	GBStandardsTabDefault      bool           `xml:"GBStandardsTabDefault,attr"`
	ReportingPeriods           []ReportPeriod `xml:"ReportingPeriods>ReportPeriod"`
	ReportingPeriod            struct {
		GradePeriod string `xml:"GradePeriod,attr"`
		StartDate   Time   `xml:"StartDate,attr"`
		EndDate     Time   `xml:"EndDate,attr"`
	} `xml:"ReportingPeriod"`
	Courses []Course `xml:"Courses>Course"`
}

type ReportPeriod struct {
	GradePeriod string `xml:"GradePeriod,attr"`
	StartDate   Time   `xml:"StartDate,attr"`
	EndDate     Time   `xml:"EndDate,attr"`
}

type Course struct {
	UsesRichContent                         bool   `xml:"UsesRichContent,attr"`
	Period                                  int8   `xml:"Period,attr"`
	Title                                   string `xml:"Title,attr"`
	Room                                    string `xml:"Room,attr"`
	Staff                                   string `xml:"Staff,attr"`
	StaffEMail                              string `xml:"StaffEMail,attr"`
	StaffGU                                 string `xml:"StaffGU,attr"`
	HighlightPercentageCutOffForProgressBar int    `xml:"HighlightPercentageCutOffForProgressBar,attr"`
	Marks                                   []Mark `xml:"Marks>Mark"`
}

type Mark struct {
	MarkName                string       `xml:"MarkName,attr"`
	CalculatedScoreString   string       `xml:"CalculatedScoreString,attr"`
	CalculatedScoreRaw      int8         `xml:"CalculatedScoreRaw,attr"`
	StandardViews           string       `xml:"StandardViews"`
	GradeCalculationSummary string       `xml:"GradeCalculationSummary"`
	Assignments             []Assignment `xml:"Assignments>Assignment"`
}

type Assignment struct {
	GradebookID        int    `xml:"GradebookID,attr"`
	Measure            string `xml:"Measure,attr"`
	Type               string `xml:"Type,attr"`
	Date               Time   `xml:"Date,attr"`
	DueDate            Time   `xml:"DueDate,attr"`
	Score              Score  `xml:"Score,attr"`
	ScoreType          string `xml:"ScoreType,attr"`
	Points             Points `xml:"Points,attr"`
	Notes              string `xml:"Notes,attr"`
	TeacherID          int    `xml:"TeacherID,attr"`
	StudentID          int    `xml:"StudentID,attr"`
	MeasureDescription string `xml:"MeasureDescription,attr"`
	HasDropBox         bool   `xml:"HasDropBox,attr"`
	DropStartDate      Time   `xml:"DropStartDate,attr"`
	DropEndDate        Time   `xml:"DropEndDate,attr"`
	Resources          string `xml:"Resources"`
	Standards          string `xml:"Standards"`
}

func (client *Client) GradeBook(builder *ParamaterBuilder) (*GradeBook, error) {
	paramater := builder.Build()
	header := DefaultHeader()
	data, err := client.Request(PXPEndpoint, PXPWebServices, GradeBookMethod, &header, &paramater)
	if err != nil {
		return nil, err
	}

	text, err := GetXmlString(*data)
	if err != nil {
		return nil, err
	}

	gb := GradeBook{}
	err = xml.Unmarshal([]byte(*text), &gb)

	if err != nil {
		return nil, err
	}

	return &gb, nil
}

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
