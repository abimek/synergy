package studentvue

const Endpoint = "/Service/PXPCommunication.asmx/ProcessWebServiceRequest"

type Identifier int

type Client struct {
	url       string
	identifer Identifier
	password  string
}

func (p Client) New(url string, identifier int, password string) Client {
	return Client{url + Endpoint, Identifier(identifier), password}
}

func (p *Client) RequestReportCard(period ReportPeriod) {
	if period.getData() == ReportPeriodNone {
		// TODO: Send Request To The Servers And Return Data
		return
	}
	// TODO: Send Request For Specific Period and Return Data
}
