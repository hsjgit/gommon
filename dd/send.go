package dd

import (
	"errors"
	"net/http"
	"strings"

	"github.com/hsjgit/gommon/httputil"
	"github.com/sirupsen/logrus"
)

type Hook struct {
	ddNotify bool
	url      string
	levels   []logrus.Level
}

func (h *Hook) Fire(e *logrus.Entry) error {
	return nil
}

type Message interface {
	Send(url string) (interface{}, error)
}

func (h *Hook) Levels() []logrus.Level {
	level := make([]logrus.Level, len(h.levels))
	for i := range h.levels {
		level = append(level, h.levels[i])
	}
	return level
}

type At struct {
	AtMobiles []string `json:"atMobiles"`
	AtUserIds []int64  `json:"atUserIds"`
	IsAtAll   bool     `json:"isAtAll"`
}

func SendMessage(ddurl, message string) error {
	if ddurl == "" {
		return errors.New("钉钉 webhook url 为空")
	}
	var headers = http.Header{}
	headers.Add("Content-Type", "application/json")
	_, err := httputil.HttpPost(ddurl, strings.NewReader(message), nil, headers)
	if err != nil {
		return err
	}
	return nil
}
