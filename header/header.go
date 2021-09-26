package header

// Header structure used when you define a custom header
// Inside your request
type Header struct {
	UserAgent string
}

//Creates a default header settings
func DefaultHeader() *Header {
	header := Header{
		UserAgent: "Walle",
	}

	return &header
}

// Receives a User-Agent string and change it in Headers
func (h *Header) SetUserAgent(newUserAgent string) {
	h.UserAgent = newUserAgent
}

// Receives a header as Header and return it as map[string][]string
func (h *Header) ParseHeader() map[string][]string {
	parsedHeader := map[string][]string{
		"User-Agent": {h.UserAgent},
	}

	return parsedHeader
}
