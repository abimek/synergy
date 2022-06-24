package studentinfo

import (
	"encoding/xml"

	"studentvue"
)

type StudentInfo struct {
	LockerInfo struct {
		LockerGU           string `xml:"LockerGU,attr"`
		LockerNumber       string `xml:"LockerNumber,attr"`
		CurrentCombination string `xml:"CurrentCombination,attr"`
		Location           string `xml:"Location,attr"`
	} `xml:"LockerInfoRecords>StudentLockerInfoRecord"`
	FormattedName      string             `xml:"FormattedName"`
	PermID             int                `xml:"PermID"`
	Gender             string             `xml:"Gender"`
	Grade              byte               `xml:"Grade"`
	Address            string             `xml:"Address"`
	LastNameGoesBy     string             `xml:"LastNameGoesBy"`
	NickName           string             `xml:"NickName"`
	BirthDate          studentvue.Time    `xml:"BirthDate"`
	EMail              string             `xml:"EMail"`
	Phone              string             `xml:"Phone"`
	HomeLanguage       string             `xml:"HomeLanguage"`
	CurrentSchool      string             `xml:"CurrentSchool"`
	Track              string             `xml:"Track"`
	HomeRoomTch        string             `xml:"HomeRoomTch"`
	HomeRoomTchEMail   string             `xml:"HomeRoomTchEMail"`
	HomeRoomTchStaffGU string             `xml:"HomeRoomTchStaffGU"`
	OrgYearGU          string             `xml:"OrgYearGU"`
	HomeRoom           string             `xml:"HomeRoom"`
	CounselorName      string             `xml:"CounselorName"`
	CounselorEmail     string             `xml:"CounselorEmail"`
	CounselorStaffGU   string             `xml:"CounselorStaffGU"`
	Photo              string             `xml:"Photo"`
	EmergencyContacts  []EmergencyContact `xml:"EmergencyContacts>EmergencyContact"`
	Physician          struct {
		Name     string `xml:"Name,attr"`
		Hospital string `xml:"Hospital,attr"`
		Phone    string `xml:"Phone,attr"`
		Extn     string `xml:"Extn,attr"`
	} `xml:"Physician"`
	Dentist struct {
		Name   string `xml:"Name,attr"`
		Office string `xml:"Office,attr"`
		Phone  string `xml:"Phone,attr"`
		Extn   string `xml:"Extn,attr"`
	} `xml:"Dentist"`
	UserDefinedGroupBoxes []UserDefinedGroupBox `xml:"UserDefinedGroupBoxes>UserDefinedGroupBox"`
}

type EmergencyContact struct {
	Name         string `xml:"Name,attr"`
	Relationship string `xml:"Relationship,attr"`
	HomePhone    string `xml:"HomePhone,attr"`
	WorkPhone    string `xml:"WorkPhone,attr"`
	OtherPhone   string `xml:"OtherPhone,attr"`
	MobilePhone  string `xml:"MobilePhone,attr"`
}

type UserDefinedGroupBox struct {
	GroupBoxLabel    string            `xml:"GroupBoxLabel,attr"`
	GroupBoxID       string            `xml:"GroupBoxID,attr"`
	VCID             string            `xml:"VCID,attr"`
	UserDefinedItems []UserDefinedItem `xml:"UserDefinedItems>UserDefinedItem"`
}

type UserDefinedItem struct {
	ItemLabel     string `xml:"ItemLabel,attr"`
	ItemType      string `xml:"ItemType,attr"`
	SourceObject  string `xml:"SourceObject,attr"`
	SourceElement string `xml:"SourceElement,attr"`
	VCID          string `xml:"VCID,attr"`
	Value         string `xml:"Value,attr"`
}

func New(client *studentvue.Client) (*StudentInfo, error) {
	params := studentvue.GetEmptyParamater()

	header := studentvue.DefaultHeader()
	data, err := client.Request(studentvue.PXPEndpoint, studentvue.PXPWebServices, studentvue.StudentInfo, &header, &params)
	if err != nil {
		return nil, err
	}

	text, err := studentvue.GetXmlString(*data)
	if err != nil {
		return nil, err
	}

	si := StudentInfo{}
	err = xml.Unmarshal([]byte(*text), &si)

	if err != nil {
		return nil, err
	}
	return &si, nil
}
