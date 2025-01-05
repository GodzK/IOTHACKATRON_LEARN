package mqtt

// Topic represents an MQTT topic
type Topic string

const (
	TopicLightSwitch Topic = "home/light/light-switch"
	TopicTempSensor  Topic = "home/sensor/temp-sensor"
)
