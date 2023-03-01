package dd

import (
	"testing"
)

func TestTextSend(t *testing.T) {
	text := NewTextMessage()
	text.Text = Text{
		Content: "业务报警, @XXX 是不一样的烟火",
	}
	text.At = at
	send, err := text.Send(dd)
	if err != nil {
		return
	}
	t.Log(send)
}
