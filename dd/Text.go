package dd

import (
	"errors"

	"github.com/hsjgit/gommon/httputil"
)

// 发送 text 类型的消息到钉钉
// 参数          参数类型  是否必填    说明
// msgtype      String   是         消息类型，此时固定为：text。
//
// content      String   是         消息内容。
//
// atMobiles    Array    否         被@人的手机号。（注意 在text内容里要有@人的手机号，只有在群内的成员才可被@，非群内成员手机号会被脱敏。）
//
// atUserIds    Array    否         被@人的用户userid。（在content里添加@人的userid。）
//
// isAtAll      Boolean  否         是否@所有人。

type textMessage struct {
	Msgtype string      `json:"msgtype"`
	Text    interface{} `json:"text"`
	At      At          `json:"at"`
}

type Text struct {
	Content string `json:"content"`
}

func (t *textMessage) Send(url string) (interface{}, error) {
	if url == "" {
		return nil, errors.New("钉钉 webhook url 为空")
	}
	var res interface{}
	err := httputil.JsonPost(url, t, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewTextMessage() textMessage {
	return textMessage{
		Msgtype: "text",
	}
}
