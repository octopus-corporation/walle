package response

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	StatusCode int
	Content    string
	Header     http.Header
}

/* Receives the raw response and converts it to a Response object
   With only StatusCode, Content and Header */
func ParseResponse(resp *http.Response) Response {
	parsedResp := Response{
		StatusCode: resp.StatusCode,
		Content:    getResponseText(resp.Body),
		Header:     resp.Header,
	}

	return parsedResp
}

/* Receives the response body and returns its data as a string */
func getResponseText(body io.ReadCloser) string {
	responseData, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}

	responseText := string(responseData)
	return responseText
}
