package studentvue

import (
	"strconv"

	"studentvue/paramaters"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
)

type Handle string

const (
	PXPWebServices Handle = "PXPWebServices"
	HDInfoServices Handle = "HDInfoServices"
)

const Endpoint = "/Service/PXPCommunication.asmx/ProcessWebServiceRequest"

type Header map[string]string

func DefaultHeader() Header {
	return map[string]string{
		"Content-Type":    "application/x-www-form-urlencoded",
		"Accept-Encoding": "gzip",
		"User-Agent":      "ksoap2-android/2.6.0+",
	}
}

func (h *Header) AddHeader(key, value string) {
	(*h)[key] = value
}

func (h *Header) ApplyHeader(r *gentleman.Request) {
	for k, v := range *h {
		r.SetHeader(k, v)
	}
}

type Identifier int

type Client struct {
	client    *gentleman.Client
	url       string
	identifer Identifier
	password  string
}

func New(url string, identifier int, password string) Client {
	url = url + Endpoint
	client := gentleman.New()
	client.URL(url)
	return Client{client, url, Identifier(identifier), password}
}

func (c *Client) request(handle *Handle, method *paramaters.Method, head *Header, paramaters *paramaters.Paramater) (*string, error) {
	request := c.client.Request()
	head.ApplyHeader(request)
	rbody := map[string]string{
		"userID":               strconv.Itoa(int(c.identifer)),
		"password":             c.password,
		"skipLoginLog":         "true",
		"parent":               "false",
		"webServiceHandleName": string(*handle),
		"methodName":           string(*method),
		"paramStr":             string(*paramaters),
	}
	request.Use(body.JSON(rbody))

	resp, ok := request.Send()

	if ok != nil {
		return nil, ok
	}

	stringVal := resp.String()
	return &stringVal, nil
}

func defaultHeaders(r *gentleman.Request) {
	r.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	r.SetHeader("Accept-Encoding", "gzip")
	r.SetHeader("User-Agent", "ksoap2-android/2.6.0+")
}
