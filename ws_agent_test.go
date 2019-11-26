package okex

import (
	"fmt"
	"testing"
	"time"
)

func GetOKExWss() *OKWSAgent {
	config := &Config{}

	config.Endpoint = "https://www.okex.com/"
	config.WSEndpoint = "wss://real.okex.com:8443/"
	config.ApiKey = ""
	config.SecretKey = ""
	config.Passphrase = ""
	config.TimeoutSecond = 45
	config.IsPrint = false
	config.I18n = ENGLISH

	okexWss := &OKWSAgent{}
	err := okexWss.Start(config)
	if err != nil {
		fmt.Println("okex wss connect failed", err)
	}

	return okexWss
}

func TestConnection(t *testing.T) {
	connections := GetOKExWss()

	channel := "spot/depth5"
	symbols := []string{"BTC-USDT", "ETH-BTC"}

	receivedDataCallback := func(obj interface{}) error {
		switch obj.(type) {
		case string:
			fmt.Println("recv string", obj)
		case int:
			fmt.Println("recv int", obj)
		case *WSDepthTableResponse:
			obj, ok := obj.(*WSDepthTableResponse)
			if !ok {
				return nil
			}

			for _, event := range obj.Data {
				fmt.Println("recv depth", event.Timestamp, event.InstrumentId, event.Asks, event.Bids)
			}
		default:
			msg, err := Struct2JsonString(obj)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			fmt.Println("recv", msg)
		}
		return nil
	}

	for _, symbol := range symbols {
		err := connections.Subscribe(channel, symbol, receivedDataCallback)
		fmt.Println("subscribe", channel, symbol, err)
	}

	chain := make(chan bool)

	go func() {
		time.Sleep(time.Second * 5)
		chain <- true
	}()

	<-chain
}
