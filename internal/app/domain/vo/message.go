package vo

var MESSAGE_TO_ALL = "*"

type Message struct {
	From    string
	To      string
	Content string
}
