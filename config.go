package main

import (
	"log"
	"time"

	"github.com/caarlos0/env"
	mqttExtCfg "github.com/mannkind/paho.mqtt.golang.ext/cfg"
)

type config struct {
	MQTT           *mqttExtCfg.MQTTConfig
	Host           string        `env:"UNIFI_HOST"                 envDefault:"unifi"`
	Port           string        `env:"UNIFI_PORT"                 envDefault:"8443"`
	Site           string        `env:"UNIFI_SITE"                 envDefault:"default"`
	Username       string        `env:"UNIFI_USERNAME"             envDefault:"unifi"`
	Password       string        `env:"UNIFI_PASSWORD"             envDefault:"unifi"`
	AwayTimeout    time.Duration `env:"UNIFI_AWAYTIMEOUT"          envDefault:"5m"`
	LookupInterval time.Duration `env:"UNIFI_LOOKUPINTERVAL"       envDefault:"10s"`
	DeviceMapping  []string      `env:"UNIFI_DEVICEMAPPING"        envDefault:"11:22:33:44:55:66;MyPhone,12:23:34:45:56:67;AnotherPhone"`
}

func newConfig(mqttCfg *mqttExtCfg.MQTTConfig) *config {
	c := config{}
	c.MQTT = mqttCfg

	if c.MQTT.ClientID == "" {
		c.MQTT.ClientID = "DefaultUnifi2MqttClientID"
	}

	if c.MQTT.DiscoveryName == "" {
		c.MQTT.DiscoveryName = "unifi"
	}

	if c.MQTT.TopicPrefix == "" {
		c.MQTT.TopicPrefix = "home/unifi"
	}

	if err := env.Parse(&c); err != nil {
		log.Printf("Error unmarshaling configuration: %s", err)
	}

	redactedPassword := ""
	if len(c.MQTT.Password) > 0 || len(c.Password) > 0 {
		redactedPassword = "<REDACTED>"
	}

	log.Printf("Environmental Settings:")
	log.Printf("  * ClientID          : %s", c.MQTT.ClientID)
	log.Printf("  * Broker            : %s", c.MQTT.Broker)
	log.Printf("  * Username          : %s", c.MQTT.Username)
	log.Printf("  * Password          : %s", redactedPassword)
	log.Printf("  * Discovery         : %t", c.MQTT.Discovery)
	log.Printf("  * DiscoveryPrefix   : %s", c.MQTT.DiscoveryPrefix)
	log.Printf("  * DiscoveryName     : %s", c.MQTT.DiscoveryName)
	log.Printf("  * TopicPrefix       : %s", c.MQTT.TopicPrefix)
	log.Printf("  * AwayTimeout       : %s", c.AwayTimeout)
	log.Printf("  * LookupInterval    : %s", c.LookupInterval)
	log.Printf("  * DeviceMapping     : %s", c.DeviceMapping)
	log.Printf("  * Host              : %s", c.Host)
	log.Printf("  * Port              : %s", c.Port)
	log.Printf("  * Username          : %s", c.Username)
	log.Printf("  * Password          : %s", redactedPassword)

	return &c
}
