package attendance

import (
	"fmt"

	"studentvue"
)

type Districts struct{}

// TODO: change return from string to Districts
func New(client *studentvue.Client) (*string, error) {
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
	fmt.Println(*text)
	return text, nil
}
