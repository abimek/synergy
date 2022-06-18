package studentvue

import "studentvue/paramaters"

// TODO: This will actually represent the data gathered
type GradeBook struct {
	data string
}

func NewGradeBook(client *Client, period *paramaters.ReportPeriod) (GradeBook, error) {
	if period.Period == paramaters.ReportPeriodNone {
		// TODO: Send Request
		return _, nil
	}

	// TODO: Send request with the specific period
	return _, nil
}
