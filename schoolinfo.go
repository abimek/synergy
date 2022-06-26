package studentvue

import (
	"encoding/xml"
)

type SchoolInfo struct {
	School         string  `xml:"School,attr"`
	Principal      string  `xml:"Principal,attr"`
	SchoolAddress  string  `xml:"SchoolAddress,attr"`
	SchoolAddress2 string  `xml:"SchoolAddress2,attr"`
	SchoolCity     string  `xml:"SchoolCity,attr"`
	SchoolState    string  `xml:"SchoolState,attr"`
	SchoolZip      int     `xml:"SchoolZip,attr"`
	Phone          string  `xml:"Phone,attr"`
	Phone2         string  `xml:"Phone2,attr"`
	URL            string  `xml:"URL,attr"`
	PrincipalEmail string  `xml:"PrincipalEmail,attr"`
	PrincipalGu    string  `xml:"PrincipalGu,attr"`
	StaffList      []Staff `xml:"StaffLists>StaffList"`
}

type Staff struct {
	Name    string `xml:"Name,attr"`
	EMail   string `xml:"EMail,attr"`
	Title   string `xml:"Title,attr"`
	Phone   string `xml:"Phone,attr"`
	Extn    string `xml:"Extn,attr"`
	StaffGU string `xml:"StaffGU,attr"`
}

func (client *Client) SchoolInfo() (*SchoolInfo, error) {
	params := GetEmptyParamater()

	header := DefaultHeader()

	data, err := client.Request(PXPEndpoint, PXPWebServices, StudentSchoolInfoMethod, &header, &params)
	if err != nil {
		return nil, err
	}

	text, err := GetXmlString(*data)
	if err != nil {
		return nil, err
	}
	si := SchoolInfo{}
	err = xml.Unmarshal([]byte(*text), &si)

	if err != nil {
		return nil, err
	}
	return &si, nil
}
