package synergy

import (
	"encoding/xml"
)

type CalendarListing struct {
	SchoolBegDate Time    `xml:"SchoolBegDate,attr"`
	SchoolEndDate Time    `xml:"SchoolEndDate,attr"`
	MonthBegDate  Time    `xml:"MonthBegDate,attr"`
	MonthEndDate  Time    `xml:"MonthEndDate,attr"`
	EventList     []Event `xml:"EventLists>EventList"`
}

// Specific Event on the calendar
type Event struct {
	Date           Time   `xml:"Date,attr"`
	Title          string `xml:"Title,attr"`
	DayType        string `xml:"DayType,attr"`
	StartTime      string `xml:"StartTime,attr"`
	Icon           string `xml:"Icon,attr"`
	AGU            string `xml:"AGU,attr"`
	Link           string `xml:"Link,attr"`
	DGU            string `xml:"DGU,attr"`
	ViewType       int    `xml:"ViewType,attr"`
	AddLinkData    string `xml:"AddLinkData,attr"`
	EvtDescription string `xml:"EvtDescription,attr"`
}

func (client *Client) CalendarListing() (*CalendarListing, error) {
	params := GetEmptyParamater()

	header := DefaultHeader()
	data, err := client.Request(PXPEndpoint, PXPWebServices, StudentCalendarMethod, &header, &params)
	if err != nil {
		return nil, err
	}

	text, err := GetXmlString(*data)
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
