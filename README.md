# OKEx Open API V3 SDK (Golang Version)

## 项目来源

https://github.com/okcoin-okex/open-api-v3-sdk

| 联系我们 | Contact Us |                    |
| -------- | ---------- | ------------------ |
| 微信号   | WeChat ID  | ApiSupport         |
| 邮 箱    | E-mail     | wei.cao@okcoin.net |

本项目去除了一些其他语言, 修改一些 bug.

## 1. 下载依赖

```bash
go get -u github.com/justinchou/okex-go-sdk-api
```

## 2. 编写 golang 文件

测试文件必须以 `\*\_test.go` 结尾, 例如: okex_open_api_v3_test.go

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

func getOKExConfig() okex.Config {
	var config okex.Config

	config.Endpoint = "https://www.okex.com/"
	config.WSEndpoint = "wss://real.okex.com:8443/"

	config.ApiKey = ""
	config.SecretKey = ""
	config.Passphrase = ""
	config.TimeoutSecond = 45
	config.IsPrint = true
	config.I18n = okex.ENGLISH

	return config
}

// 创建 https 客户端
func NewOKExClient() *okex.Client {
	config := getOKExConfig()
	client := okex.NewClient(config)
	return client
}

// 创建 wss 客户端
func NewOKExWss() (client *okex.OKWSAgent, err error) {
	config := getOKExConfig()

	client := &okex.OKWSAgent{}
	err = client.Start(config)
	if err != nil {
		fmt.Println("okex wss connect failed", err)
		return nil
	}

	return client
}
```

## 3. 启动测试

```bash
go test -v -run TestOKExServerTime okex_open_api_v3_test.go
```
