package client

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Client 		==> 客户端实例
type Client struct {
	Request *Request
	Cookie  *http.Cookie
	Result  Result
}

// Request 		==> 请求体
type Request struct {
	Url           string
	Method        string
	Data          io.Reader
	ContentType   string
	Authorization string
}

// Result 		==> 结果集
type Result struct {
	Body   []byte
	Status int
}

// Do 		==> 执行请求
func (c *Client) Do() {
	//HTTP请求构造
	request, _ := http.NewRequest(c.Request.Method, c.Request.Url, c.Request.Data)
	request.Header.Set("Content-Type", c.Request.ContentType)
	if c.Request.Authorization != "" {
		request.Header.Set("Authorization", c.Request.Authorization)
	}
	if c.Cookie.String() != "" {
		request.AddCookie(c.Cookie)
	}
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(res.Cookies()) > 1 {
		c.Cookie = res.Cookies()[1]
	}
	defer res.Body.Close()
	c.Result.Status = res.StatusCode
	c.Result.Body, _ = ioutil.ReadAll(res.Body)

}
