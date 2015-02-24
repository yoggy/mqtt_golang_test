package main

import (
	"./config"
	"fmt"
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"time"
)

func main() {
	conf := config.GetConfig()

	opts := MQTT.NewClientOptions()
	opts.AddBroker(conf.BrokerURL)
	opts.SetClientId("mqtt_pub.go")
	if conf.UseAuth {
		opts.SetUsername(conf.Username)
		opts.SetPassword(conf.Password)
	}

	client := MQTT.NewClient(opts)
	if _, err := client.Start(); err != nil {
		panic(err)
	}

	for i := 0; i < 1000; i++ {
		text := fmt.Sprintf("message no=%d", i)
		fmt.Printf("publish : topic=%s, payload=%s\n", conf.Topic, text)
		receipt := client.Publish(MQTT.QoS(0), conf.Topic, []byte(text))
		<-receipt
		time.Sleep(1 * time.Second)
	}
}
