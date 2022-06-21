package studentvue

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Handle string

const (
	PXPWebServices Handle = "PXPWebServices"
	HDInfoServices Handle = "HDInfoServices"
)

const Endpoint = "/Service/PXPCommunication.asmx/ProcessWebServiceRequest"

type Header map[string]string

func DefaultHeader() Header {
	return Header(map[string]string{
		"Content-Type":    "application/x-www-form-urlencoded",
		"Accept-Encoding": "gzip",
		"User-Agent":      "ksoap2-android/2.6.0+",
	})
}

func (h *Header) AddHeader(key, value string) {
	(*h)[key] = value
}

func (h *Header) ApplyHeader(r *http.Request) {
	for k, v := range *h {
		r.Header.Set(k, v)
	}
}

type Identifier int

type Client struct {
	client    *http.Client
	url       string
	identifer Identifier
	password  string
}

func New(url string, identifier int, password string) Client {
	url = url + Endpoint
	client := &http.Client{}
	return Client{client, url, Identifier(identifier), password}
}

func (c *Client) Request(handle Handle, method Method, head *Header, paramaters *Paramater) (*string, error) {
	data := url.Values{}
	data.Set("userID", strconv.Itoa(int(c.identifer)))
	data.Set("password", c.password)
	data.Set("skipLoginLog", "true")
	data.Set("parent", "false")
	data.Set("webServiceHandleName", string(handle))
	data.Set("methodName", string(method))
	data.Set("paramStr", string(*paramaters))
	req, _ := http.NewRequest("POST", c.url, strings.NewReader(data.Encode()))
	head.ApplyHeader(req)
	res, ok := c.client.Do(req)
	if ok != nil {
		return nil, ok
	}
	stringVa, _ := ioutil.ReadAll(res.Body)
	stringVal := string(stringVa)
	return &stringVal, nil
}
