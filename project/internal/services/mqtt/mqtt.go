package mqtt

import (
	"fmt"

	pahoMqtt "github.com/eclipse/paho.mqtt.golang"
)

type Mqtt struct {
	client pahoMqtt.Client
}

// NewMqtt initializes a new MQTT client
func NewMqtt() *Mqtt {
	cfg := newMqttCfg()

	// Create MQTT client options
	ops := pahoMqtt.NewClientOptions()
	ops.AddBroker(fmt.Sprintf("tcp://%s:%d", cfg.Host, cfg.Port))
	ops.SetClientID(cfg.ClientID)
	ops.SetUsername(cfg.Username)
	ops.SetPassword(cfg.Password)

	// Create and connect the MQTT client
	client := pahoMqtt.NewClient(ops)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return &Mqtt{
		client: client,
	}
}

// Subscribe subscribes to the given topic
func (m *Mqtt) Subscribe(topic Topic, callback pahoMqtt.MessageHandler) error {
	if token := m.client.Subscribe(string(topic), 0, callback); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

// Publish publishes a message to the given topic
func (m *Mqtt) Publish(topic Topic, payload string) error {
	if token := m.client.Publish(string(topic), 0, false, payload); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
