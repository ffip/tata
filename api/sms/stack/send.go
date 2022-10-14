package stack

import (
	"fmt"

	"github.com/ffip/tata/lib/net/http/client"
	"github.com/ffip/tata/lib/text"
	"github.com/tidwall/gjson"
)

// Send ==> send message to phone.
func (m *Message) Send() (err error) {
	uri := text.Mgr("https://api.1cloudsp.com/api/v2/single_send?accesskey=", _accessKey, "&secret=", _secret, "&sign=", _sign, "&templateId=", m.TemplateID, "&mobile=", m.Phone, "&content=", m.Context)
	http := client.Client{Request: &client.Request{Url: uri}}
	http.Do()

	if gjson.GetBytes(http.Result.Body, "code").Str != "0" {
		return fmt.Errorf("send error")
	}
	return
}
