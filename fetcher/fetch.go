package fetcher

import (
	"net/http"
	"io/ioutil"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"fmt"
)

func Fetch(url string) ([]byte,error) {

	client := &http.Client{}

	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	//reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	//reqest.Header.Add("Accept-Encoding", "gzip, deflate")
	//reqest.Header.Add("Accept-Language", "zh-cn,zh;q=0.8,en-us;q=0.5,en;q=0.3")
	//reqest.Header.Add("Connection", "keep-alive")
	//reqest.Header.Add("Host", "login.sina.com.cn")
	//reqest.Header.Add("Referer", "http://weibo.com/")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	resp, err := client.Do(reqest)

	//resp,err := http.Get(url)
	if err != nil {
		return nil , err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil , fmt.Errorf("return code is not %s",http.StatusOK)
	}
	bodyreader := bufio.NewReader(resp.Body)
	e := determineEncodeing(*bodyreader)
	utf8Reader := transform.NewReader(bodyreader,e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)

}

func determineEncodeing(r bufio.Reader) encoding.Encoding  {
	b , err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e , _ , _ := charset.DetermineEncoding(b,"")

	return e
}