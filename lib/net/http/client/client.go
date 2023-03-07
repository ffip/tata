package client

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
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
	UserAgent     string
	Header        map[string]string
	Timeout       time.Duration
	// The proxy type is determined by the URL scheme. "http",
	// "https", and "socks5" are supported. If the scheme is empty,
	//
	// If Proxy is nil or nil *URL, no proxy is used.
	ProxyUrl url.URL
}

// Result 		==> 结果集
type Result struct {
	Body   []byte
	Status int
}

// NewRequest 		==> 新建请求
func NewRequest() *Client {
	return &Client{Request: &Request{UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"}}
}

// Do 		==> 执行请求
func (c *Client) Do() *Client {
	//HTTP请求构造
	request, _ := http.NewRequest(c.Request.Method, c.Request.Url, c.Request.Data)
	request.Header.Set("Content-Type", c.Request.ContentType)
	request.Header.Set("Referer", c.Request.Url)
	if c.Request.Authorization != "" {
		request.Header.Set("Authorization", c.Request.Authorization)
	}
	if c.Request.UserAgent != "" {
		request.Header.Set("User-Agent", c.Request.UserAgent)
	}
	if c.Cookie.String() != "" {
		request.AddCookie(c.Cookie)
	}
	// 支持自定义Header
	for k, v := range c.Request.Header {
		request.Header.Set(k, v)
	}

	var client *http.Client
	if c.Request.Timeout != 0 {
		client.Timeout = c.Request.Timeout
	}
	if c.Request.ProxyUrl == (url.URL{}) {
		client = &http.Client{}
	} else {
		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(&c.Request.ProxyUrl)}}
	}
	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return c
	}
	if len(res.Cookies()) > 1 {
		c.Cookie = res.Cookies()[1]
	}
	defer res.Body.Close()
	c.Result.Status = res.StatusCode
	c.Result.Body, _ = ioutil.ReadAll(res.Body)
	return c
}

// Get 		==> 定义请求方式
func (c *Client) Get() *Client {
	c.Request.Method = "GET"
	return c
}

// Post 		==> 定义请求方式
func (c *Client) Post() *Client {
	c.Request.Method = "POST"
	return c
}

// Put 		==> 定义请求方式
func (c *Client) Put() *Client {
	c.Request.Method = "PUT"
	return c
}

// Delete 		==> 定义请求方式
func (c *Client) Delete() *Client {
	c.Request.Method = "DELETE"
	return c
}

// SetUrl 		==> 定义请求目标
func (c *Client) SetUrl(url string) *Client {
	c.Request.Url = url
	return c
}

// SetMethod 		==> 定义请求方法
func (c *Client) SetMethod(method string) *Client {
	c.Request.Method = method
	return c
}

// SetContentType 		==> 定义内容类型
func (c *Client) SetContentType(contentType string) *Client {
	c.Request.ContentType = contentType
	return c
}

// SetBody 		==> 定义请求内容
func (c *Client) SetBody(body io.Reader) *Client {
	c.Request.Data = body
	return c
}

// SetAuthorization 		==> 定义身份验证
func (c *Client) SetAuthorization(credentials string) *Client {
	c.Request.Authorization = credentials
	return c
}

// GetStatusCode 		==> 获取请求状态码
func (c *Client) GetStatusCode() int {
	return c.Result.Status
}

// GetBody 		==> 获取返回内容
func (c *Client) GetBody() []byte {
	return c.Result.Body
}

// GetBody 		==> 获取返回内容
func (c *Client) GetBodyString() string {
	return string(c.Result.Body)
}

// SaveToFile 		==> 写出结果到文件
func (c *Client) SaveToFile(filepath string) error {
	// Create the download file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	err = ioutil.WriteFile(filepath, c.GetBody(), 0777)
	if err != nil {
		return err
	}
	return nil
}
