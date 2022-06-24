package calendar

import (
	"encoding/xml"

	"studentvue"
)

type CalendarListing struct {
	SchoolBegDate studentvue.Time `xml:"SchoolBegDate,attr"`
	SchoolEndDate studentvue.Time `xml:"SchoolEndDate,attr"`
	MonthBegDate  studentvue.Time `xml:"MonthBegDate,attr"`
	MonthEndDate  studentvue.Time `xml:"MonthEndDate,attr"`
	EventList     []Event         `xml:"EventLists>EventList"`
}

// Specific Event on the calendar
type Event struct {
	Date           studentvue.Time `xml:"Date,attr"`
	Title          string          `xml:"Title,attr"`
	DayType        string          `xml:"DayType,attr"`
	StartTime      string          `xml:"StartTime,attr"`
	Icon           string          `xml:"Icon,attr"`
	AGU            string          `xml:"AGU,attr"`
	Link           string          `xml:"Link,attr"`
	DGU            string          `xml:"DGU,attr"`
	ViewType       int             `xml:"ViewType,attr"`
	AddLinkData    string          `xml:"AddLinkData,attr"`
	EvtDescription string          `xml:"EvtDescription,attr"`
}

func New(client *studentvue.Client) (*CalendarListing, error) {
	params := studentvue.GetEmptyParamater()

	header := studentvue.DefaultHeader()
	data, err := client.Request(studentvue.PXPEndpoint, studentvue.PXPWebServices, studentvue.StudentCalendar, &header, &params)
	if err != nil {
		return nil, err
	}

	text, err := studentvue.GetXmlString(*data)
	if err != nil {
		return nil, err
	}

	cl := CalendarListing{}

	err = xml.Unmarshal([]byte(*text), &cl)

	if err != nil {
		return nil, err
	}
	return &cl, err
}
