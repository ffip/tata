package stack

import (
	"github.com/ffip/tata/lib/net/http/client"
	"github.com/ffip/tata/lib/text"
	"github.com/tidwall/gjson"
)

// Message ==> send message api.
type Message struct {
	TemplateID string
	Phone      string
	Context    string
}

const (
	_accessKey = "dxw786601305"
	_secret    = "XiwwSfShSTgcNsPPg5t6bTQByVAp7kAz"
	_sign      = "【民杰教育】"
)

// Send ==> send message to phone.
func (m *Message) Send() {
	uri := text.Mgr("https://api.1cloudsp.com/api/v2/single_send?accesskey=", _accessKey, "&secret=", _secret, "&sign=", _sign, "&templateId=", m.TemplateID, "&mobile=", m.Phone, "&content=", m.Context)
	http := client.Client{Request: &client.Request{Url: uri}}
	http.Do()

	if gjson.GetBytes(http.Result.Body, "code").Str != "0" {
		return
	}

	// 将验证短信发送记录存储到数据库
	return
}
