package gradebook

import (
	"encoding/xml"

	"studentvue"
)

type GradeBook struct {
	Type                       string `xml:"Type,attr"`
	ErrorMessage               string `xml:"ErrorMessage,attr"`
	HideStandardGraphInd       bool   `xml:"HideStandardGraphInd,attr"`
	HideMarksColumnElementary  bool   `xml:"HideMarksColumnElementary,attr"`
	HidePointsColumnElementary bool   `xml:"HidePointsColumnElementary,attr"`
	HidePercentSecondary       bool   `xml:"HidePercentSecondary,attr"`
	DisplayStandardsData       bool   `xml:"DisplayStandardsData,attr"`
	GBStandardsTabDefault      bool   `xml:"GBStandardsTabDefault,attr"`
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
							Points             Points          `xml:"Points,attr"`
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

// NewGradeBook will either return a GradeBook or an error
//
func New(client *studentvue.Client, builder *studentvue.ParamaterBuilder) (*GradeBook, error) {
	paramater := builder.Build()
	header := studentvue.DefaultHeader()
	data, err := client.Request(studentvue.PXPWebServices, studentvue.GradeBook, &header, &paramater)
	if err != nil {
		return nil, err
	}

	text, err := studentvue.GetXmlString(*data)
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
