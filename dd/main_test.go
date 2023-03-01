package dd

import (
	"testing"
)

var at At

const dd = "https://oapi.dingtalk.com/robot/send?access_token=63fd584a8c9db8d8d60c1a71b244931e5e2ba0553a7a93d32e76bc1685a79a44"

func TestMain(m *testing.M) {
	at = At{
		AtMobiles: []string{"150XXXXXXXX"},
		AtUserIds: []int64{21312},
		IsAtAll:   true,
	}
	m.Run()

}
