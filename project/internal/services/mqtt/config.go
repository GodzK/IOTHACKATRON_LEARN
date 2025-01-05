package mqtt

import (
	"github.com/Netflix/go-env"
)

// MqttCfg holds the MQTT configuration
type MqttCfg struct {
	Host     string `env:"MQTT_HOST,default=localhost,required=true"`
	Port     int    `env:"MQTT_PORT,default=1883,required=true"`
	Username string `env:"MQTT_USERNAME,required=true"`
	Password string `env:"MQTT_PASSWORD,required=true"`
	ClientID string `env:"MQTT_CLIENT_ID,required=true"`
}

func newMqttCfg() *MqttCfg {
	var cfg MqttCfg
	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}
