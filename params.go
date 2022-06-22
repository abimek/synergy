package studentvue

import "strconv"

type Method string

type Paramater string

const (
	GradeBook               Method = "Gradebook"
	GetMatchingDistrictList Method = "GetMatchingDistrictList"
	Attendance              Method = "Attendance"
)

const (
	ReportPeriodNone = -1
)

type ParamaterBuilder struct {
	paramaters string
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
	return "<MatchToDistrictZipCode>" + strconv.Itoa(p.ZipCode) + "</MatchToDistrictZipCode>"
}
