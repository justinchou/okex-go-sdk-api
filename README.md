# OKEx Open API V3 SDK (Golang Version)

## 项目来源

https://github.com/okcoin-okex/open-api-v3-sdk

| 联系我们 | Contact Us |                    |
| -------- | ---------- | ------------------ |
| 微信号   | WeChat ID  | ApiSupport         |
| 邮 箱    | E-mail     | wei.cao@okcoin.net |

本项目去除了一些其他语言, 修改一些 bug.

## 1.Downloads or updates OKEX code's dependencies, in your command line

```bash
go get -u github.com/justinchou/okex-go-sdk-api
```

## 2.Write the go file. warm tips: test go file, must suffix \*\_test.go, eg: okex_open_api_v3_test.go

```go
package gotest

import (
	"fmt"
	"github.com/justinchou/okex-go-sdk-api"
	"testing"
)

func TestOKExServerTime(t *testing.T) {
	serverTime, err := NewOKExClient().GetServerTime()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("OKEx's server time: ", serverTime)
}

func NewOKExClient() *okex.Client {
	var config okex.Config

	config.Endpoint = "https://www.okex.com/"
	config.WSEndpoint = "wss://real.okex.com:8443/"

	config.ApiKey = ""
	config.SecretKey = ""
	config.Passphrase = ""
	config.TimeoutSecond = 45
	config.IsPrint = true
	config.I18n = okex.ENGLISH

	client := okex.NewClient(config)
	return client
}
```

## 3. run test go

```bash
go test -v -run TestOKExServerTime okex_open_api_v3_test.go
```
