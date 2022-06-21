package gradebook

import (
	"encoding/xml"
	"fmt"
	"strings"

	"studentvue"
	"studentvue/paramaters"
)

type GradeBook struct {
	XMLName                    xml.Name `xml:"Gradebook"`
	Xsd                        string   `xml:"xsd,attr"`
	Xsi                        string   `xml:"xsi,attr"`
	Type                       string   `xml:"Type,attr"`
	ErrorMessage               string   `xml:"ErrorMessage,attr"`
	HideStandardGraphInd       bool     `xml:"HideStandardGraphInd,attr"`
	HideMarksColumnElementary  bool     `xml:"HideMarksColumnElementary,attr"`
	HidePointsColumnElementary bool     `xml:"HidePointsColumnElementary,attr"`
	HidePercentSecondary       bool     `xml:"HidePercentSecondary,attr"`
	DisplayStandardsData       bool     `xml:"DisplayStandardsData,attr"`
	GBStandardsTabDefault      bool     `xml:"GBStandardsTabDefault,attr"`
	ReportingPeriods           struct {
		ReportPeriod []struct {
			GradePeriod string          `xml:"GradePeriod,attr"`
			StartDate   studentvue.Time `xml:"StartDate,attr"`
			EndDate     studentvue.Time `xml:"EndDate,attr"`
		} `xml:"ReportPeriod"`
	} `xml:"ReportingPeriods"`
	ReportingPeriod struct {
		GradePeriod string          `xml:"GradePeriod,attr"`
		StartDate   studentvue.Time `xml:"StartDate,attr"`
		EndDate     studentvue.Time `xml:"EndDate,attr"`
	} `xml:"ReportingPeriod"`
	Courses struct {
		Course []struct {
			UsesRichContent                         bool   `xml:"UsesRichContent,attr"`
			Period                                  int8   `xml:"Period,attr"`
			Title                                   string `xml:"Title,attr"`
			Room                                    string `xml:"Room,attr"`
			Staff                                   string `xml:"Staff,attr"`
			StaffEMail                              string `xml:"StaffEMail,attr"`
			StaffGU                                 string `xml:"StaffGU,attr"`
			HighlightPercentageCutOffForProgressBar int    `xml:"HighlightPercentageCutOffForProgressBar,attr"`
			Marks                                   struct {
				Mark []struct {
					MarkName                string `xml:"MarkName,attr"`
					CalculatedScoreString   string `xml:"CalculatedScoreString,attr"`
					CalculatedScoreRaw      int8   `xml:"CalculatedScoreRaw,attr"`
					StandardViews           string `xml:"StandardViews"`
					GradeCalculationSummary string `xml:"GradeCalculationSummary"`
					Assignments             struct {
						Assignment []struct {
							GradebookID        int             `xml:"GradebookID,attr"`
							Measure            string          `xml:"Measure,attr"`
							Type               string          `xml:"Type,attr"`
							Date               studentvue.Time `xml:"Date,attr"`
							DueDate            studentvue.Time `xml:"DueDate,attr"`
							Score              Score           `xml:"Score,attr"`
							ScoreType          string          `xml:"ScoreType,attr"`
							Points             string          `xml:"Points,attr"`
							Notes              string          `xml:"Notes,attr"`
							TeacherID          int             `xml:"TeacherID,attr"`
							StudentID          int             `xml:"StudentID,attr"`
							MeasureDescription string          `xml:"MeasureDescription,attr"`
							HasDropBox         bool            `xml:"HasDropBox,attr"`
							DropStartDate      studentvue.Time `xml:"DropStartDate,attr"`
							DropEndDate        studentvue.Time `xml:"DropEndDate,attr"`
							Resources          string          `xml:"Resources"`
							Standards          string          `xml:"Standards"`
						} `xml:"Assignment"`
					} `xml:"Assignments"`
				} `xml:"Mark"`
			} `xml:"Marks"`
		} `xml:"Course"`
	} `xml:"Courses"`
}

type GradeType string

const (
	Normal          GradeType = "normal"
	NotGraded       GradeType = "not graded"
	CharachterGrade GradeType = "charachter grade"
)

type book struct {
	XMLName xml.Name `xml:"Book"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
}

func NewGradeBook(client *studentvue.Client, period *paramaters.ReportPeriod) (*GradeBook, error) {
	var paramater paramaters.Paramater
	builder := paramaters.NewParamaterBuilder()
	if period.Period == paramaters.ReportPeriodNone {
		paramater = builder.Build()
	} else {
		builder.Add(period)
		paramater = builder.Build()
	}
	header := studentvue.DefaultHeader()
	data, ok := client.Request(studentvue.PXPWebServices, paramaters.GradeBook, &header, &paramater)

	if ok != nil {
		return nil, ok
	}

	sxml := strings.Replace(*data, "string", "Book", 2)

	b := book{}
	err := xml.Unmarshal([]byte(sxml), &b)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	gb := GradeBook{}
	err = xml.Unmarshal([]byte(b.Text), &gb)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &gb, nil
}
