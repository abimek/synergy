package reportcard

import (
	"encoding/xml"
	"errors"

	"github.com/abimekuriya/studentvue"
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

// I don't currently know how to decode the Base64Code into a PDF
type ReportCard struct {
	DocumentGU  string `xml:"DocumentGU,attr"`
	FileName    string `xml:"FileName,attr"`
	DocFileName string `xml:"DocFileName,attr"`
	DocType     string `xml:"DocType,attr"`
	Base64Code  string `xml:"Base64Code"`
}

// Returns a ReportingPeriods struct which contains all the reporting periods
func New(client *studentvue.Client) (*ReportingPeriods, error) {
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

func (p *ReportingPeriod) GetReportCard(client *studentvue.Client) (*ReportCard, error) {
	if p.DocumentGU == "" {
		return nil, errors.New("DocumentGU is null")
	}
	builder := studentvue.NewParamaterBuilder()
	builder.Add(&studentvue.DocumentGU{DocumentGU: p.DocumentGU})

	param := builder.Build()
	header := studentvue.DefaultHeader()
	data, err := client.Request(studentvue.PXPEndpoint, studentvue.PXPWebServices, studentvue.GetReportCardDocumentData, &header, &param)
	if err != nil {
		return nil, err
	}
	text, err := studentvue.GetXmlString(*data)
	if err != nil {
		return nil, err
	}
	rc := ReportCard{}
	err = xml.Unmarshal([]byte(*text), &rc)
	if err != nil {
		return nil, err
	}
	return &rc, nil
}
