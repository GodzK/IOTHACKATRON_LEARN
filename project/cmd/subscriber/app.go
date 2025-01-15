package main

import (
	"log"
	"server/internal/connectors"
	"server/internal/repositories/ultrasonic"
	"server/internal/services/mqtt"
	"strconv"
	"time"

	pahoMqtt "github.com/eclipse/paho.mqtt.golang"
)

type App struct {
	mqtt           *mqtt.Mqtt
	ultrasonicRepo *ultrasonic.UltrasonicRepo
}

func NewApp() *App {
	mqtt := mqtt.NewMqtt()
	db := connectors.NewDatabase()
	ultrasonicRepo := ultrasonic.NewUltrasonicRepo(db.DB)
	return &App{
		mqtt:           mqtt,
		ultrasonicRepo: ultrasonicRepo,
	}
}

func (a *App) Run() {
	// Subscribe to the temp sensor topic
	for {
		// Asynchronous callback for subscribing to the temperature sensor
		a.mqtt.Subscribe(mqtt.TopicTempSensor, func(client pahoMqtt.Client, msg pahoMqtt.Message) {
			content := string(msg.Payload())
			// Attempt to convert the message content to an integer
			number, err := strconv.Atoi(content)
			if err != nil {
				log.Println("Error converting payload to int:", err)
				return
			}

			// Create a new ultrasonic object
			ultrasonic := ultrasonic.Ultrasonic{
				Value: float64(number),

				DateTimestamp: time.Now().Format("2006-01-02 15:04:05"),
			}

			// Insert the ultrasonic data into the database
			if err := a.ultrasonicRepo.Insert(ultrasonic); err != nil {
				log.Println("Error inserting ultrasonic data:", err)
				log.Println("Error inserting ultrasonic data:", err)
				return
			}

			// Check if the number is greater than 30 and publish a message to the light switch topic
			if number > 30 {
				a.mqtt.Publish(mqtt.TopicLightSwitch, "ON")
			} else {
				a.mqtt.Publish(mqtt.TopicLightSwitch, "OFF")
			}
		})
		// Add a break or timeout mechanism if needed to stop the infinite loop (for graceful shutdown)
	}
}
