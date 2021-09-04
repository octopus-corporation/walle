package requests

import (
	"net/http"

	"github.com/octopus-corporation/walle/response"
)

type Request struct {
	Url    string
	Method string
}

type Callback struct {
	OnError            func(error)
	OnResponseReceived func(response.Response)
}

/*Receives a request object and a callback object.
  Checks the request method and calls the appropriate function.
  Returns a response with StatusCode, Headers and Content*/
func Enter(request Request, callback Callback) {
	var resp *http.Response
	var err error

	switch request.Method {
	case "GET":
		resp, err = get(request.Url)
	}

	if err != nil {
		callback.OnError(err)
		return
	}

	defer resp.Body.Close()

	parsedResp := response.ParseResponse(resp)
	callback.OnResponseReceived(parsedResp)
}

/* Does a get request to the url and returns the response. */
func get(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	return resp, err
}
