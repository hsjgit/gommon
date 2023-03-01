package dd

import (
	"errors"

	"github.com/hsjgit/gommon/httputil"
)

// 发送 markdown 类型的消息到钉钉
// 参数          参数类型  是否必填    说明
// msgtype      String   是         消息类型，此时固定为：markdown。
//
// title        String   是         首屏会话透出的展示内容。
//
// text         String   是         markdown格式的消息。
//
// atMobiles    Array    否         被@人的手机号。（注意 在text内容里要有@人的手机号，只有在群内的成员才可被@，非群内成员手机号会被脱敏。）

// atUserIds    Array    否         被@人的用户userid。（在content里添加@人的userid。）
//
// isAtAll      Boolean  否         是否@所有人。

type markdownMessage struct {
	Msgtype  string      `json:"msgtype"`
	Markdown interface{} `json:"markdown"`
	At       At          `json:"at"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func (m *markdownMessage) Send(url string) (interface{}, error) {
	if url == "" {
		return nil, errors.New("钉钉 webhook url 为空")
	}
	var res interface{}
	err := httputil.JsonPost(url, m, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewMarkdownMessage() markdownMessage {
	return markdownMessage{
		Msgtype: "markdown",
	}
}
