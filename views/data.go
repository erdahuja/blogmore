package views

const (
	AlertLvlError   = "danger"
	AlertLvlSuccess = "success"
	AlertMsgGeneric = "Something went wrong"
)

// Alert model represents general purpose alert box below header
// in templates
type Alert struct {
	Level   string
	Message string
}

// Data represents top level structure that views can expect
type Data struct {
	Alert *Alert
	Yield  interface{}
}
