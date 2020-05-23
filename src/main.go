package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	"gomongo.mqtt/src/models"
	mqtt "gomongo.mqtt/src/protocol"
)

func main() {
	uri, err := url.Parse("tcp://localhost:1880")
	if err != nil {
		log.Fatal(err)
	}
	topic := "data"
	go mqtt.Listen(uri, topic)

	client := mqtt.Connect("pub", uri)

	timer := time.NewTicker(1 * time.Second)
	for t := range timer.C {
		t.String()
		sensorData := &models.SensorData{
			TimeAquisition: time.Now(),
			Value:          20.0,
		}
		b, err := json.Marshal(sensorData)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Publicando no topico '%s'", topic)
		client.Publish(topic, 0, false, b)
	}
}
