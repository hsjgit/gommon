package dd

import (
	"runtime"
	"testing"
)

func TestA(t *testing.T) {
	stack := make([]byte, 1024*8)
	stack = stack[:runtime.Stack(stack, false)]
	t.Log(string(stack))

}

func TestMarkdownSend(t *testing.T) {
	m := NewMarkdownMessage()
	m.At = at
	m.Markdown = Markdown{
		Title: "业务报警",
		Text:  "#### 杭州天气 @150XXXXXXXX \n > 9度，西北风1级，空气良89，相对温度73%\n > ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n > ###### 10点20分发布 [天气](https://www.dingtalk.com) \n",
	}
	send, err := m.Send(dd)
	if err != nil {
		return
	}
	t.Log(send)
}
