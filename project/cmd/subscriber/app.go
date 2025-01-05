package main

import (
	"fmt"
	"server/internal/services/mqtt"

	pahoMqtt "github.com/eclipse/paho.mqtt.golang"
)

type App struct {
	mqtt *mqtt.Mqtt
}

func NewApp() *App {
	mqtt := mqtt.NewMqtt()
	return &App{
		mqtt: mqtt,
	}
}
func (a *App) Run() {
	// Subscribe to the temp sensor topic
	for {
		a.mqtt.Subscribe(mqtt.TopicTempSensor, func(client pahoMqtt.Client, msg pahoMqtt.Message) {
			content := string(msg.Payload())
			// Print the message content
			fmt.Println(content)
		})
	}
}
