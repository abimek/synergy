package attendance

import (
	"encoding/xml"

	"studentvue"
)

type Attendance struct {
	Type        string `xml:"Type,attr"`
	StartPeriod byte   `xml:"StartPeriod,attr"`
	EndPeriod   byte   `xml:"EndPeriod,attr"`
	PeriodCount byte   `xml:"PeriodCount,attr"`
	SchoolName  string `xml:"SchoolName,attr"`
	Absences    struct {
		Absence []struct {
			AbsenceDate           studentvue.Time `xml:"AbsenceDate,attr"`
			Reason                string          `xml:"Reason,attr"`
			Note                  string          `xml:"Note,attr"`
			DailyIconName         string          `xml:"DailyIconName,attr"`
			CodeAllDayReasonType  string          `xml:"CodeAllDayReasonType,attr"`
			CodeAllDayDescription string          `xml:"CodeAllDayDescription,attr"`
			Periods               struct {
				Period []struct {
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
				} `xml:"Period"`
			} `xml:"Periods"`
		} `xml:"Absence"`
	} `xml:"Absences"`
	TotalExcused struct {
		PeriodTotal []PeriodTotal `xml:"PeriodTotal"`
	} `xml:"TotalExcused"`
	TotalTardies struct {
		PeriodTotal []PeriodTotal `xml:"PeriodTotal"`
	} `xml:"TotalTardies"`
	TotalUnexcused struct {
		PeriodTotal []PeriodTotal `xml:"PeriodTotal"`
	} `xml:"TotalUnexcused"`
	TotalActivities struct {
		PeriodTotal []PeriodTotal `xml:"PeriodTotal"`
	} `xml:"TotalActivities"`
	TotalUnexcusedTardies struct {
		PeriodTotal []PeriodTotal `xml:"PeriodTotal"`
	} `xml:"TotalUnexcusedTardies"`
	ConcurrentSchoolsLists struct {
		ConcurrentSchoolsList struct {
			ConcurrentOrgYearGU  string `xml:"ConcurrentOrgYearGU,attr"`
			ConcurrentSchoolName string `xml:"ConcurrentSchoolName,attr"`
		} `xml:"ConcurrentSchoolsList"`
	} `xml:"ConcurrentSchoolsLists"`
}

type PeriodTotal struct {
	Number byte `xml:"Number,attr"`
	Total  byte `xml:"Total,attr"`
}

// TODO: change return from string to Districts
func New(client *studentvue.Client) (*Attendance, error) {
	builder := studentvue.NewParamaterBuilder()
	params := builder.Build()
	header := studentvue.DefaultHeader()
	data, err := client.Request(studentvue.PXPWebServices, studentvue.Attendance, &header, &params)
	if err != nil {
		return nil, err
	}

	text, err := studentvue.GetXmlString(*data)
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
