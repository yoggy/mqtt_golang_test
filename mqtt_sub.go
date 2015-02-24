package main

import (
	"./config"
	"fmt"
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

var messageHandler MQTT.MessageHandler = func(client *MQTT.MqttClient, msg MQTT.Message) {
	fmt.Printf("message received : topic=%s, payload=%s\n", msg.Topic(), string(msg.Payload()))
}

func main() {
	conf := config.GetConfig()

	opts := MQTT.NewClientOptions()
	opts.AddBroker(conf.BrokerURL)
	opts.SetDefaultPublishHandler(messageHandler)
	opts.SetClientId("mqtt_sub.go")
	if conf.UseAuth {
		opts.SetUsername(conf.Username)
		opts.SetPassword(conf.Password)
	}

	client := MQTT.NewClient(opts)
	if _, err := client.Start(); err != nil {
		panic(err)
	}

	// subscribe topic
	filter, _ := MQTT.NewTopicFilter(conf.Topic, 0)
	if receipt, err := client.StartSubscription(nil, filter); err != nil {
		panic(err)
	} else {
		<-receipt // wait for StartSubscription
	}

	// wait forever...
	select {}
}
