/*
 * @author      Liu Yongshuai<liuyongshuai@hotmail.com>
 * @package     http
 * @date        2018-01-25 19:19
 */
package http

import (
	"net/http"
	"strings"
)

//响应结构体，在response基础上封装
type ToFuResponse struct {
	body             []byte            //响应的body
	status           string            //响应码的描述信息，"200 OK"
	statusCode       int               //响应码"200"
	proto            string            //所用协议"HTTP/1.1"
	header           map[string]string //头信息
	contentLen       int64             //返回内容长度
	setCookie        []http.Cookie     //要设置的cookie信息
	location         string            //当statusCode为3XX如301时重定向的链接
	err              error             //请求的出错信息
	transferEncoding []string          //所用的编码信息
}

//提取body信息
func (tfr ToFuResponse) GetBody() []byte {
	return tfr.body
}

//返回响应的body的字符串格式
func (tfr ToFuResponse) GetBodyString() string {
	return string(tfr.body)
}

//提取状态描述信息，如"200 OK"
func (tfr ToFuResponse) GetStatus() string {
	return tfr.status
}

//提取状态码，如200
func (tfr ToFuResponse) GetStatusCode() int {
	return tfr.statusCode
}

//所用协议"HTTP/1.1"
func (tfr ToFuResponse) GetProto() string {
	return tfr.proto
}

//提取响应的头信息
func (tfr ToFuResponse) GetHeader() map[string]string {
	return tfr.header
}

//获取响应信息的长度
func (tfr ToFuResponse) GetContentLen() int64 {
	return tfr.contentLen
}

//获取响应头里面set-cookie设置cookie信息的列表
func (tfr ToFuResponse) GetSetCookie() []http.Cookie {
	return tfr.setCookie
}

//获取重定向后的地址信息
func (tfr ToFuResponse) GetLocation() string {
	return tfr.location
}

//提取响应的编码信息
func (tfr ToFuResponse) GetTransferEncoding() []string {
	return tfr.transferEncoding
}

//返回出错的信息
func (tfr ToFuResponse) Error() error {
	return tfr.err
}

//对原始cookie进行五马分尸
func SplitRawCookie(ck string) (ret map[string]string) {
	ck = strings.TrimSpace(ck)
	if len(ck) == 0 {
		return
	}
	kvs := strings.Split(ck, ";")
	if len(kvs) == 0 {
		return
	}
	for _, val := range kvs {
		val = strings.TrimSpace(val)
		if !strings.Contains(val, "=") {
			continue
		}
		ind := strings.Index(val, "=")
		k := strings.TrimSpace(val[0:ind])
		v := strings.TrimSpace(val[ind+1:])
		if len(k) == 0 || len(v) == 0 {
			continue
		}
		ret[k] = v
	}
	return
}

//合并cookie
func JoinRawCookie(ck map[string]string) (ret string) {
	if len(ret) == 0 {
		return ""
	}
	var tmp []string
	for k, v := range ck {
		tmp = append(tmp, k+"="+v)
	}
	ret = strings.Join(tmp, "; ")
	return
}
