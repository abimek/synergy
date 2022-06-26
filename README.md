# studentvue

Structured API calls to EDUPoint's StudentVue alongside representitive data structs.

## Disclaimer
**The API as of now is still being worked on and is flexible to drastic
change and should not be used until it is stabilized**

## Installation
```bash
go get github.com/abimek/synergy
```

## Features
- GradeBook
- Attendance
- ReportCard
- Calendar
- StudentInfo
- SchoolInfo

## ToDo
- [ ] Implement Proper Districting
- [ ] Student Class List
- [ ] Studnet Health Info
- [ ] Login Confirmation

## Example

```go
package main

import (
	"fmt"

	studentvue "github.com/abimek/synergy"
)

func main() {
	client := studentvue.New("school portal", 0o000000, "password")

	pb := studentvue.ParamaterBuilder{}
	pb.Add(&studentvue.ReportPeriodParamater{Period: 0})

	gradebook, err := client.GradeBook(&pb)
	if err != nil {
		fmt.Println("issue getting grade")
	}

	// Print the points gained on the fourth assignment in marking period 1 in course 1
	fmt.Println(gradebook.Courses[0].Marks[0].Assignments[3].Score.Points)
}
```
