package http

import (
	"testing"
	"fmt"
)

var testUrl = "https://www.baidu.com/s?wd=golang%20openfile&pn=610&oq=golang%20openfile&ie=utf-8&rsv_idx=1&rsv_pq=9535c6380000b84b&rsv_t=e7eaF5wQgnDd0ZUG50OodOQsa8zV%2BW5LFT%2FZ0Mt3vt00Lt6DHSEcBW7wsp0"
var ua = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36`
var rawCookie = "BAIDUID=394004972595DCD4F12ED179531FB596:FG=1; BIDUPSID=394004972595DCD4F12ED179531FB596; PSTM=1516244001; BD_UPN=123253; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; H_PS_PSSID=1447_21118_20930; BD_CK_SAM=1; PSINO=2; BD_HOME=0; H_PS_645EC=56121cJQbMsCUEK6llQKRqY65SybhlglilOT2tgKPeEMvOmntngo1UjcKf8; BDSVRTM=103"

func TestToFuHttp(t *testing.T) {
	fmt.Println(t.Name())
	size := 20
	ch := make(chan int, size)
	for i := 0; i < size; i++ {
		go fortest(ch)
	}
	for i := 0; i < size; i++ {
		<-ch
	}
}

func fortest(ch chan int) {
	client := NewHttpClient().
		SetUrl(testUrl).
		SetUserAgent(ua).
		SetTimeout(30).
		SetKeepAlive(true).
		SetRawCookie(rawCookie).
		AddHeader("X-Requested-With", "XMLHttpRequest").
		AddHeader("Connection", "keep-alive")
	ret, err := client.Get()
	fmt.Println(err, ret.GetBodyString())
	ch <- 1
}
