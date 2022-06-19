package studentvue

import "studentvue/paramaters"

// TODO: This will actually represent the data gathered
type GradeBook struct {
	data string
}

func NewGradeBook(client *Client, period *paramaters.ReportPeriod) (*GradeBook, error) {
	var paramater paramaters.Paramater
	builder := paramaters.NewParamaterBuilder()
	if period.Period == paramaters.ReportPeriodNone {
		paramater = builder.Build()
	} else {
		builder.AddParamater(period)
		paramater = builder.Build()
	}
	header := DefaultHeader()
	data, ok := client.request(PXPWebServices, paramaters.GradeBook, &header, &paramater)

	if ok != nil {
		return nil, ok
	}
	gradebook := GradeBook{*data}
	return &gradebook, nil
}
