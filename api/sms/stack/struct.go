package stack

// Message ==> send message api.
type Message struct {
	TemplateID string
	Phone      string
	Context    string
}

const (
	_accessKey = "key"
	_secret    = "value"
	_sign      = "【test】"
)
