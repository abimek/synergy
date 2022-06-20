package studentvue

import (
	"encoding/xml"
	"fmt"
	"strings"

	"studentvue/paramaters"
)

type GradeBook struct {
	XMLName                    xml.Name `xml:"Gradebook"`
	Xsd                        string   `xml:"xsd,attr"`
	Xsi                        string   `xml:"xsi,attr"`
	Type                       string   `xml:"Type,attr"`
	ErrorMessage               string   `xml:"ErrorMessage,attr"`
	HideStandardGraphInd       string   `xml:"HideStandardGraphInd,attr"`
	HideMarksColumnElementary  string   `xml:"HideMarksColumnElementary,attr"`
	HidePointsColumnElementary string   `xml:"HidePointsColumnElementary,attr"`
	HidePercentSecondary       string   `xml:"HidePercentSecondary,attr"`
	DisplayStandardsData       string   `xml:"DisplayStandardsData,attr"`
	GBStandardsTabDefault      string   `xml:"GBStandardsTabDefault,attr"`
	ReportingPeriods           struct {
		ReportPeriod []struct {
			Text        string `xml:",chardata"`
			Index       string `xml:"Index,attr"`
			GradePeriod string `xml:"GradePeriod,attr"`
			StartDate   string `xml:"StartDate,attr"`
			EndDate     string `xml:"EndDate,attr"`
		} `xml:"ReportPeriod"`
	} `xml:"ReportingPeriods"`
	ReportingPeriod struct {
		GradePeriod string `xml:"GradePeriod,attr"`
		StartDate   string `xml:"StartDate,attr"`
		EndDate     string `xml:"EndDate,attr"`
	} `xml:"ReportingPeriod"`
	Courses struct {
		Course []struct {
			UsesRichContent                         string `xml:"UsesRichContent,attr"`
			Period                                  string `xml:"Period,attr"`
			Title                                   string `xml:"Title,attr"`
			Room                                    string `xml:"Room,attr"`
			Staff                                   string `xml:"Staff,attr"`
			StaffEMail                              string `xml:"StaffEMail,attr"`
			StaffGU                                 string `xml:"StaffGU,attr"`
			HighlightPercentageCutOffForProgressBar string `xml:"HighlightPercentageCutOffForProgressBar,attr"`
			Marks                                   struct {
				Mark struct {
					MarkName                string `xml:"MarkName,attr"`
					CalculatedScoreString   string `xml:"CalculatedScoreString,attr"`
					CalculatedScoreRaw      string `xml:"CalculatedScoreRaw,attr"`
					StandardViews           string `xml:"StandardViews"`
					GradeCalculationSummary string `xml:"GradeCalculationSummary"`
					Assignments             struct {
						Assignment []struct {
							GradebookID        string `xml:"GradebookID,attr"`
							Measure            string `xml:"Measure,attr"`
							Type               string `xml:"Type,attr"`
							Date               string `xml:"Date,attr"`
							DueDate            string `xml:"DueDate,attr"`
							Score              string `xml:"Score,attr"`
							ScoreType          string `xml:"ScoreType,attr"`
							Points             string `xml:"Points,attr"`
							Notes              string `xml:"Notes,attr"`
							TeacherID          string `xml:"TeacherID,attr"`
							StudentID          string `xml:"StudentID,attr"`
							MeasureDescription string `xml:"MeasureDescription,attr"`
							HasDropBox         string `xml:"HasDropBox,attr"`
							DropStartDate      string `xml:"DropStartDate,attr"`
							DropEndDate        string `xml:"DropEndDate,attr"`
							Resources          string `xml:"Resources"`
							Standards          string `xml:"Standards"`
						} `xml:"Assignment"`
					} `xml:"Assignments"`
				} `xml:"Mark"`
			} `xml:"Marks"`
		} `xml:"Course"`
	} `xml:"Courses"`
}

type Book struct {
	XMLName xml.Name `xml:"Book"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
}

func NewGradeBook(client *Client, period *paramaters.ReportPeriod) (*GradeBook, error) {
	// TODO: Proper Error Handeling

	var paramater paramaters.Paramater
	builder := paramaters.NewParamaterBuilder()
	if period.Period == paramaters.ReportPeriodNone {
		paramater = builder.Build()
	} else {
		builder.Add(period)
		paramater = builder.Build()
	}
	header := DefaultHeader()
	data, ok := client.request(PXPWebServices, paramaters.GradeBook, &header, &paramater)

	if ok != nil {
		return nil, ok
	}

	sxml := strings.Replace(*data, "string", "Book", 2)

	b := Book{}
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
