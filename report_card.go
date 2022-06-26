package studentvue

import (
	"encoding/xml"
	"errors"
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
func (client *Client) ReportCard() (*ReportingPeriods, error) {
	params := GetEmptyParamater()
	header := DefaultHeader()

	data, err := client.Request(PXPEndpoint, PXPWebServices, GetReportCardInitialDataMethod, &header, &params)
	if err != nil {
		return nil, err
	}
	text, err := GetXmlString(*data)
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

func (p *ReportingPeriod) GetReportCard(client *Client) (*ReportCard, error) {
	if p.DocumentGU == "" {
		return nil, errors.New("DocumentGU is null")
	}
	builder := NewParamaterBuilder()
	builder.Add(&DocumentGUParmater{DocumentGU: p.DocumentGU})

	param := builder.Build()
	header := DefaultHeader()
	data, err := client.Request(PXPEndpoint, PXPWebServices, GetReportCardDocumentDataMethod, &header, &param)
	if err != nil {
		return nil, err
	}
	text, err := GetXmlString(*data)
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
