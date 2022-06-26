package synergy

import (
	"encoding/xml"
)

type Attendance struct {
	Type                   string             `xml:"Type,attr"`
	StartPeriod            byte               `xml:"StartPeriod,attr"`
	EndPeriod              byte               `xml:"EndPeriod,attr"`
	PeriodCount            byte               `xml:"PeriodCount,attr"`
	SchoolName             string             `xml:"SchoolName,attr"`
	Absences               []Absence          `xml:"Absences>Absence"`
	TotalExcused           []PeriodTotal      `xml:"TotalExcused>PeriodTotal"`
	TotalTardies           []PeriodTotal      `xml:"TotalTardies>PeriodTotal"`
	TotalUnexcused         []PeriodTotal      `xml:"TotalUnexcused>PeriodTotal"`
	TotalActivities        []PeriodTotal      `xml:"TotalActivities>PeriodTotal"`
	TotalUnexcusedTardies  []PeriodTotal      `xml:"TotalUnexcusedTardies"`
	ConcurrentSchoolsLists []ConcurrentSchool `xml:"ConcurrentSchoolsLists>ConcurrentSchoolsList"`
}

// A signle instance / day which you were absence
type Absence struct {
	AbsenceDate           Time     `xml:"AbsenceDate,attr"`
	Reason                string   `xml:"Reason,attr"`
	Note                  string   `xml:"Note,attr"`
	DailyIconName         string   `xml:"DailyIconName,attr"`
	CodeAllDayReasonType  string   `xml:"CodeAllDayReasonType,attr"`
	CodeAllDayDescription string   `xml:"CodeAllDayDescription,attr"`
	Periods               []Period `xml:"Periods>Period"`
}

// This type represents a signle Period in which you were absent
type Period struct {
	Number     byte   `xml:"Number,attr"`
	Name       string `xml:"Name,attr"`
	Reason     string `xml:"Reason,attr"`
	Course     string `xml:"Course,attr"`
	Staff      string `xml:"Staff,attr"`
	StaffEMail string `xml:"StaffEMail,attr"`
	IconName   string `xml:"IconName,attr"`
	SchoolName string `xml:"SchoolName,attr"`
	StaffGU    string `xml:"StaffGU,attr"`
	OrgYearGU  string `xml:"OrgYearGU,attr"`
}

// This represents a specific concurrent school
type ConcurrentSchool struct {
	ConcurrentOrgYearGU  string `xml:"ConcurrentOrgYearGU,attr"`
	ConcurrentSchoolName string `xml:"ConcurrentSchoolName,attr"`
}

// Represents one speciifc Period
type PeriodTotal struct {
	Number byte `xml:"Number,attr"`
	Total  byte `xml:"Total,attr"`
}

func (client *Client) Attendance() (*Attendance, error) {
	params := GetEmptyParamater()
	header := DefaultHeader()
	data, err := client.Request(PXPEndpoint, PXPWebServices, AttendanceMethod, &header, &params)
	if err != nil {
		return nil, err
	}

	text, err := GetXmlString(*data)
	if err != nil {
		return nil, err
	}

	at := Attendance{}
	err = xml.Unmarshal([]byte(*text), &at)

	if err != nil {
		return nil, err
	}

	return &at, err
}
