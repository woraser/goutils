/*
 * @author      Liu Yongshuai<liuyongshuai@hotmail.com>
 * @package     http
 * @date        2018-01-25 19:19
 */
package http

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

//客户端结构体
type ToFuHttp struct {
	buf     *bytes.Buffer     //发送的数据，一般POST用
	vals    url.Values        //提交上来的数据
	writer  *multipart.Writer //POST/PUT时的写入数据
	client  *http.Client      //连接客户端
	request *http.Request     //要发送的请求
}

//实例化一个端
func NewHttpClient() *ToFuHttp {
	b := new(bytes.Buffer)
	httpReq := &ToFuHttp{
		buf:    b,
		vals:   make(url.Values),
		writer: multipart.NewWriter(b),
	}
	transport := http.Transport{
		DisableKeepAlives: false,
	}
	httpReq.client = &http.Client{
		Timeout:   time.Duration(int64(30) * int64(time.Second)),
		Transport: &transport,
	}
	httpReq.request = &http.Request{}
	httpReq.request.Header = make(http.Header)
	httpReq.AddHeaders(map[string]string{
		"Cache-Control": "max-age=0",
		"User-Agent":    "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.79 Safari/537.36",
	})
	return httpReq
}

//添加要上传的文件
func (httpReq *ToFuHttp) AddFile(fieldName, filePath, fileName string) error {
	formFile, err := httpReq.writer.CreateFormFile(fieldName, fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Create form file failed: %s", fileName)
		return err
	}
	srcFile, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open file failed: %s", fileName)
		return err
	}
	defer srcFile.Close()
	_, err = io.Copy(formFile, srcFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "copy file failed: %s", fileName)
		return err
	}
	return nil
}

//添加header信息
func (httpReq *ToFuHttp) AddHeader(k string, v string) *ToFuHttp {
	httpReq.request.Header.Add(k, v)
	return httpReq
}

//批量设置头信息
func (httpReq *ToFuHttp) AddHeaders(hs map[string]string) *ToFuHttp {
	for k, v := range hs {
		httpReq.AddHeader(k, v)
	}
	return httpReq
}

//设置要请求的host（设置header的相应值）
func (httpReq *ToFuHttp) SetHost(host string) *ToFuHttp {
	httpReq.AddHeader("Host", host)
	return httpReq
}

//设置URL
func (httpReq *ToFuHttp) SetUrl(u string) *ToFuHttp {
	tu, _ := url.Parse(u)
	httpReq.request.URL = tu
	return httpReq
}

//设置长连接选项
func (httpReq *ToFuHttp) SetKeepAlive(b bool) *ToFuHttp {
	httpReq.client.Transport.(*http.Transport).DisableKeepAlives = b
	return httpReq
}

//设置userAgent（设置header的相应值）
func (httpReq *ToFuHttp) SetUserAgent(ua string) *ToFuHttp {
	httpReq.AddHeader("User-Agent", ua)
	return httpReq
}

//设置cookie（设置header的相应值）
func (httpReq *ToFuHttp) SetRawCookie(ck string) *ToFuHttp {
	httpReq.AddHeader("Cookie", ck)
	return httpReq
}

//添加单个cookie的键值
func (httpReq *ToFuHttp) AddCookie(k, v string) *ToFuHttp {
	httpReq.AddCookies(map[string]string{k: v})
	return httpReq
}

//批量添加cookie的键值
func (httpReq *ToFuHttp) AddCookies(ck map[string]string) *ToFuHttp {
	if len(ck) == 0 {
		return httpReq
	}
	cks := httpReq.request.Header.Get("Cookie")
	if len(cks) == 0 {
		cks = httpReq.request.Header.Get("cookie")
	}
	kvs := SplitRawCookie(cks)
	for k, v := range ck {
		k, v = strings.TrimSpace(k), strings.TrimSpace(v)
		if len(k) <= 0 || len(v) <= 0 {
			continue
		}
		kvs[k] = v
	}
	rck := JoinRawCookie(kvs)
	httpReq.SetRawCookie(rck)
	return httpReq
}

//设置referer（设置header的相应值）
func (httpReq *ToFuHttp) SetReferer(referer string) *ToFuHttp {
	httpReq.AddHeader("Referer", referer)
	return httpReq
}

//设置超时时间（设置header的相应值）
func (httpReq *ToFuHttp) SetTimeout(timeout int) *ToFuHttp {
	t := int64(timeout) * int64(time.Second)
	httpReq.client.Timeout = time.Duration(t)
	return httpReq
}

//设置代理用的和端口（设置header的相应值）
func (httpReq *ToFuHttp) SetProxy(proxyHost string) *ToFuHttp {
	//如果只给了IP:PORT这样的，默认为http方式
	check, _ := regexp.MatchString(`^[\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}:[\d]{1,5}$`, proxyHost)
	if check {
		proxyHost = "http://" + proxyHost
	}
	httpReq.client.Transport.(*http.Transport).Proxy = func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxyHost)
	}
	return httpReq
}

//批量添加字段
func (httpReq *ToFuHttp) AddFields(data map[string]string) *ToFuHttp {
	if len(data) > 0 {
		for k, v := range data {
			httpReq.vals.Set(k, v)
		}
	}
	return httpReq
}

//添加单个字段
func (httpReq *ToFuHttp) AddField(k, v string) *ToFuHttp {
	httpReq.vals.Set(k, v)
	return httpReq
}

//便于额外的写入请求的数据，获取writer
func (httpReq *ToFuHttp) GetWriter() *multipart.Writer {
	return httpReq.writer
}

//发起GET请求并返回数据
func (httpReq *ToFuHttp) Get() (ToFuResponse, error) {
	httpReq.request.Method = http.MethodGet
	httpReq.writer.Close()
	defer httpReq.buf.Reset()
	response, err := httpReq.client.Do(httpReq.request)
	if err != nil {
		return ToFuResponse{}, err
	}
	return processResponse(response)
}

//发起POST请求并返回数据
func (httpReq *ToFuHttp) PostBin() (ToFuResponse, error) {
	httpReq.request.Method = http.MethodPost
	u := httpReq.request.URL.String()
	httpReq.writer.Close()
	response, err := httpReq.client.PostForm(u, httpReq.vals)
	if err != nil {
		fmt.Println(err)
		return ToFuResponse{}, err
	}
	return processResponse(response)
}

//发起POST请求并返回数据
func (httpReq *ToFuHttp) Post() (ToFuResponse, error) {
	httpReq.request.Method = http.MethodPost
	for k, vs := range httpReq.vals {
		if len(vs) <= 0 {
			continue
		}
		httpReq.writer.WriteField(k, vs[0])
	}
	httpReq.writer.Close()
	defer httpReq.buf.Reset()
	//拼装请求的body
	ct := httpReq.writer.FormDataContentType()
	bf := httpReq.buf
	u := httpReq.request.URL.String()
	response, err := httpReq.client.Post(u, ct, bf)
	if err != nil {
		fmt.Println(err)
		return ToFuResponse{}, err
	}
	return processResponse(response)
}

//处理响应信息
func processResponse(response *http.Response) (ToFuResponse, error) {
	ret := ToFuResponse{header: make(map[string]string)}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ret.err = err
		return ret, err
	}
	ret.body = body
	for k, vs := range response.Header {
		ret.header[k] = strings.Join(vs, ";")
	}
	for _, cptr := range response.Cookies() {
		ret.setCookie = append(ret.setCookie, *cptr)
	}
	ret.transferEncoding = response.TransferEncoding
	ret.status = response.Status
	ret.statusCode = response.StatusCode
	ret.proto = response.Proto
	ret.contentLen = response.ContentLength
	loc, err := response.Location()
	if err == nil {
		ret.location = loc.String()
	}
	return ret, nil
}
