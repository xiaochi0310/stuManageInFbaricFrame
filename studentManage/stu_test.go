package studentManage

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

/**
 * @Author: WuNaiChi
 * @Date: 2020/7/9 10:54
 * @Desc:
 */
//const URL = "http://192.168.9.188:8090/%s"
const URL = "http://127.0.0.1:8090/%s"
const (
	GET  = "GET"
	POST = "POST"
)

var Admin string

func Test001(t *testing.T) {
	data := `{
		"acctId":"1000",
		"name":"wunaichi",
		"sex":"woman",
		"grade":"1",
		"hobby":"music"								
	}`
	payload := []byte(data)
	url := fmt.Sprintf(URL, "v1/students/createStu")
	err := SendHttp(payload, url, POST, Admin)
	if err != nil {
		t.Log(err)
	}
}
func Test002(t *testing.T) {
	data := `{
		"pageNo":1,
		"pageSize":10					
	}`
	payload := []byte(data)
	url := fmt.Sprintf(URL, "v1/students/")
	err := SendHttp(payload, url, POST, Admin)
	if err != nil {
		t.Log(err)
	}
}

// 发送HTTP请求
func SendHttp(payload []byte, url, method, customerId string) error {
	req, _ := http.NewRequest(method, url, strings.NewReader(string(payload)))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("customer-id", customerId)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("sentHttp error", err)
		return err
	}
	if res == nil {
		return nil
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
