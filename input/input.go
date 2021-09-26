package input

// Default definition of input
type Input interface{}

// Used for inputs that have some date or reference that
// the spider needs to access some specific data inside the
// source.
type ReferenceInput struct {
	Content      map[string]interface{}
	SentManually bool
	Priority     int
}

// Used when the input is just for activating the spider,
// and has no importance
type NullInput struct {
	SentManually bool
}
