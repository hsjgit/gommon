package dd

import (
	"errors"

	"github.com/hsjgit/gommon/httputil"
)

// 发送 link 类型的消息到钉钉
// 参数          参数类型  是否必填    说明
// msgtype      String   是         消息类型，此时固定为：link。
//
// title        String   是         消息标题。
//
// text         String   是         消息内容。如果太长只会部分展示。
//
// messageUrl   String   是         点击消息跳转的URL，打开方式如下： 移动端，在钉钉客户端内打开 PC端 默认侧边栏打开 希望在外部浏览器打开，请参考消息链接说明 https://open.dingtalk.com/document/orgapp-server/message-link-description
//
// picUrl       String   否         图片URL。

type linkMessage struct {
	Msgtype string      `json:"msgtype"`
	Link    interface{} `json:"link"`
}

type Link struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	MessageUrl string `json:"messageUrl"`
	PicUrl     string `json:"picUrl"`
}

func (l *linkMessage) Send(url string) (interface{}, error) {
	if url == "" {
		return nil, errors.New("钉钉 webhook url 为空")
	}
	var res interface{}
	err := httputil.JsonPost(url, l, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewLinkMessage() linkMessage {
	return linkMessage{
		Msgtype: "link",
	}
}
