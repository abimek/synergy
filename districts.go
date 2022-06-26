package studentvue

import (
	"fmt"
)

// TODO: change return from string to Districts
// Method will remain private until it's functional
func (client *Client) district(builder *ParamaterBuilder) (*string, error) {
	params := builder.Build()
	header := DefaultHeader()
	data, err := client.Request(HDEndpoint, HDInfoServices, GetMatchingDistrictListMethod, &header, &params)
	if err != nil {
		return nil, err
	}

	text, err := GetXmlString(*data)
	if err != nil {
		return nil, err
	}
	fmt.Println(*text)
	return text, err
}
