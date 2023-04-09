package mqtt

import (
	"com/anoop/examples/internal/models"
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type IotoMqttConnection struct {
	mqtt  mqtt.Client
	topic string
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Mqtt Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Mqtt Connection lost: %v", err)
}

func NewIotoMqttConnection(broker string, username string, password string, topic string) *IotoMqttConnection {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(fmt.Sprintf("go_mqtt_client_%s", username))
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return &IotoMqttConnection{mqtt: client, topic: topic}
}

func (m *IotoMqttConnection) Publish(message models.IotoMessage) {
	text, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	token := m.mqtt.Publish(m.topic, 0, false, text)
	token.Wait()
	time.Sleep(time.Second)
}

func (m *IotoMqttConnection) Sub(topic string) {
	token := m.mqtt.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}
