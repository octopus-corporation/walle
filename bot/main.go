/*
	This is an example of a crawler working with
	Walle's flow. In main we call the core function
	passing a callback, it will start receive messages
	from RabbitMQ. Once the input came, it checks the
	fields and start a new Request.
*/

package main

import (
	"fmt"

	"github.com/octopus-corporation/walle/input"
	"github.com/octopus-corporation/walle/response"
	"github.com/octopus-corporation/walle/walle"
)

func main() {
	walle.GetInputs(onReceiveInput)
}

func onReceiveInput(input input.Input) {
	fmt.Println("Received input")
	if !IsInputValid(input) {
		panic("Wrong input")
	}

	firstRequest(input)
}

func firstRequest(input input.Input) {
	request := CreateFirstRequest()
	request.Go()
}

func OnFirstPage(response response.Response) {
	fmt.Println("I received a response")
	fmt.Println(response)
}
