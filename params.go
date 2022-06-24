package studentvue

import "strconv"

type Method string

type Paramater string

const (
	GradeBook                 Method = "Gradebook"
	GetMatchingDistrictList   Method = "GetMatchingDistrictList"
	Attendance                Method = "Attendance"
	StudentCalendar           Method = "StudentCalendar"
	StudentInfo               Method = "StudentInfo"
	StudentSchoolInfo         Method = "StudentSchoolInfo"
	GetReportCardInitialData  Method = "GetReportCardInitialData"
	GetReportCardDocumentData Method = "GetReportCardDocumentData"
)

const (
	ReportPeriodNone = -1
)

type ParamaterBuilder struct {
	paramaters string
}

func GetEmptyParamater() Paramater {
	pb := NewParamaterBuilder()
	return pb.Build()
}

func NewParamaterBuilder() ParamaterBuilder {
	return ParamaterBuilder{""}
}

func (p *ParamaterBuilder) Add(paramater ParamaterType) {
	p.paramaters += paramater.ToString() + "\n"
}

func (p *ParamaterBuilder) Build() Paramater {
	return Paramater("<Params>\n" + p.paramaters + "</Params>")
}

// I decided to make my own interface instead of using the built in Stringer interface as it would be helpful
// to destinguish any datatype that can convert to a string to one's that are valid to this program, im not
// compeltely sure on wether this is best practice.

type ParamaterType interface {
	ToString() string
}

// GradeBook Method: Report Periods species which period to get grades from when using the GradeBook Method
type ReportPeriod struct {
	Period int
}

func (p *ReportPeriod) ToString() string {
	return "<ReportPeriod>" + strconv.Itoa(p.Period) + "</ReportPeriod>"
}

// GetMatchingDistrictList Method: Specifies a specific zip code to return all the schools in it
type MatchToDistrictZipCode struct {
	ZipCode int
}

func (p *MatchToDistrictZipCode) ToString() string {
	return "<Key>5E4B7859-B805-474B-A833-FDB15D205D40</Key>\n<MatchToDistrictZipCode>" + strconv.Itoa(p.ZipCode) + "</MatchToDistrictZipCode>"
}

// Used in the ReportCard DocumentData Method to request a specific document
type DocumentGU struct {
	DocumentGU string
}

func (p *DocumentGU) ToString() string {
	return "<DocumentGU>" + p.DocumentGU + "</DocumentGU>"
}
