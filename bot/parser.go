/*
	This is a parser file for abstratic some logics
	from the spider flow, like checking the inputs,
	creating the request body, finding some data inside
	a text, etc.
*/

package main

import (
	"github.com/octopus-corporation/walle/header"
	"github.com/octopus-corporation/walle/input"
	"github.com/octopus-corporation/walle/requests"
)

func IsInputValid(input input.Input) bool {
	return true
}

func CreateFirstRequest() requests.Request {
	firstRequest := requests.Request{
		Url:    "https://www.google.com",
		Method: "GET",
		Header: *header.DefaultHeader(),
		Callback: requests.Callback{
			OnResponseReceived: OnFirstPage,
		},
	}

	return firstRequest
}
