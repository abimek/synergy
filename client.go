package studentvue

import (
	"gopkg.in/h2non/gentleman.v2"
)

const (
	PXPWebServices = "PXPWebServices"
	HDInfoServices = "HDInfoServices"
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

func (c *Client) request(head Header) {
	request := c.client.Request()
	head.ApplyHeader(request)
}

func defaultHeaders(r *gentleman.Request) {
	r.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	r.SetHeader("Accept-Encoding", "gzip")
	r.SetHeader("User-Agent", "ksoap2-android/2.6.0+")
}
