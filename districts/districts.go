package districts

import (
	"fmt"

	"studentvue"
)

// TODO: change return from string to Districts
func New(client *studentvue.Client, builder *studentvue.ParamaterBuilder) (*string, error) {
	params := builder.Build()
	header := studentvue.DefaultHeader()
	data, err := client.Request(studentvue.HDEndpoint, studentvue.HDInfoServices, studentvue.GetMatchingDistrictList, &header, &params)
	if err != nil {
		return nil, err
	}

	text, err := studentvue.GetXmlString(*data)
	if err != nil {
		return nil, err
	}
	fmt.Println(*text)
	return text, err
}
