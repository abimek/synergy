package reportcard

import (
	"encoding/xml"
	"fmt"

	"studentvue"
)

type ReportingPeriods struct {
	ReportingPeriods []ReportingPeriod `xml:"RCReportingPeriods>RCReportingPeriod"`
}

type ReportingPeriod struct {
	ReportingPeriodGU   string `xml:"ReportingPeriodGU,attr"`
	ReportingPeriodName string `xml:"ReportingPeriodName,attr"`
	EndDate             string `xml:"EndDate,attr"`
	Message             string `xml:"Message,attr"`
	DocumentGU          string `xml:"DocumentGU,attr"`
}

// Returns a ReportingPeriods struct which contains all the reporting periods
func NewList(client *studentvue.Client) (*ReportingPeriods, error) {
	params := studentvue.GetEmptyParamater()
	header := studentvue.DefaultHeader()

	data, err := client.Request(studentvue.PXPEndpoint, studentvue.PXPWebServices, studentvue.GetReportCardInitialData, &header, &params)
	if err != nil {
		return nil, err
	}
	text, err := studentvue.GetXmlString(*data)
	if err != nil {
		return nil, err
	}
	rp := ReportingPeriods{}
	err = xml.Unmarshal([]byte(*text), &rp)
	if err != nil {
		return nil, err
	}
	return &rp, nil
}

func New(client *studentvue.Client, paramater *studentvue.Paramater) (*string, error) {
	header := studentvue.DefaultHeader()
	data, err := client.Request(studentvue.PXPEndpoint, studentvue.PXPWebServices, studentvue.GetReportCardDocumentData, &header, paramater)
	if err != nil {
		return nil, err
	}
	text, err := studentvue.GetXmlString(*data)
	if err != nil {
		return nil, err
	}
	fmt.Println(*text)
	return nil, nil
}
