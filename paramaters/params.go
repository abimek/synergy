package paramaters

import "strconv"

const (
	GradeBook = "Gradebook"
)

const (
	ReportPeriodNone = -1
)

type ParamaterBuilder struct {
	paramaters string
}

func (p *ParamaterBuilder) AddParamater(paramater ParamaterType) {
	p.paramaters += paramater.ToString() + "\n"
}

func (p *ParamaterBuilder) String() string {
	return p.paramaters
}

// I decided to make my own interface instead of using the built in Stringer interface as it would be helpful
// to destinguish any datatype that can convert to a string to one's that are valid to this program, im not
// compeltely sure on wether this is best practice.

type ParamaterType interface {
	ToString() string
}

type ReportPeriod struct {
	Period int
}

func (p *ReportPeriod) ToString() string {
	return "<ReportPeriod>" + strconv.Itoa(p.Period) + "</ReportPeriod>"
}
