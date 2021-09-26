package walle

import (
	"github.com/octopus-corporation/walle/input"
)

// Type of the function that is goign to be called
// When it received some input
type InputCallback func(input.Input)

// Consumes the RabbitMQ queue passing the inputs to the
// Callback function
func GetInputs(callback InputCallback) {
	input := input.ReferenceInput{
		Content: map[string]interface{}{
			"date": "2019-01-01",
		},
		SentManually: false,
		Priority:     1,
	}

	callback(input)
}
