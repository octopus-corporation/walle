package requests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/octopus-corporation/walle/header"
	"github.com/octopus-corporation/walle/response"
)

type Request struct {
	Url      string
	Method   string
	Body     string
	Callback Callback
	Header   header.Header
	Settings Settings
}

type Callback struct {
	OnError            func(error)
	OnResponseReceived func(response.Response)
}

type Settings struct {
	Async bool
}

/*
Receives a request object. Checks the request method
and call the appropriate function. Returns a response
with StatusCode, Headers and Content
*/
func enter(r *Request) {
	var resp *http.Response
	var err error

	switch r.Method {
	case "GET":
		resp, err = get(r.Url)
	case "POST":
		resp, err = post(r.Url, r.Header, r.Body)
	}

	if err != nil {
		r.Callback.OnError(err)
		return
	}

	defer resp.Body.Close()

	parsedResp := response.ParseResponse(resp)
	r.Callback.OnResponseReceived(parsedResp)
}

/* Does a get request to the url and returns the response. */
func get(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	return resp, err
}

/* Does a post request to the url and returns the response. */
func post(reqUrl string, reqHeader header.Header, reqBody string) (*http.Response, error) {
	parsedUrl, err := url.Parse(reqUrl)
	if err != nil {
		return nil, err
	}
	parsedHeader := reqHeader.ParseHeader()
	parsedBody := ioutil.NopCloser(strings.NewReader(reqBody))
	req := &http.Request{
		Method: "POST",
		URL:    parsedUrl,
		Header: parsedHeader,
		Body:   parsedBody,
	}

	res, err := http.DefaultClient.Do(req)
	return res, err
}

// Starts the request according to the request definition.
// Also defines if the request is going to be asynchronous
// acording to the request settings
func (r *Request) Go() {
	fmt.Println(r.Method, r.Url, "| Async", r.Settings.Async)
	if r.Settings.Async {
		go enter(r)
	} else {
		enter(r)
	}
}
