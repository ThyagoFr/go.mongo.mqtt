package mqtt

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	"gomongo.mqtt/src/models"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Connect - aaaaaaaaaaaaaaaaaaaaaaaaaaaa
func Connect(clientID string, uri *url.URL) mqtt.Client {

	options := CreateClientOptions(clientID, uri)
	client := mqtt.NewClient(options)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client

}

// CreateClientOptions - bbbbbbbbbbbbbbbbbbbbbbbbb
func CreateClientOptions(clientID string, uri *url.URL) *mqtt.ClientOptions {

	options := mqtt.NewClientOptions()
	options.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	options.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	options.SetPassword(password)
	options.SetClientID(clientID)
	return options

}

// Listen - Go lint eu ti odeio
func Listen(uri *url.URL, topic string) {
	client := Connect("sub", uri)
	fmt.Printf("Assinando topico '%s'...\n", topic)
	client.Subscribe(topic, 0, func(client mqtt.Client, message mqtt.Message) {
		dataSensor := &models.SensorData{}
		err := json.Unmarshal(message.Payload(), dataSensor)
		if err != nil {
			log.Fatal(err)
		}
		dataSensor.InsertData()
		fmt.Println("Payload", string(message.Payload()))
	})

}
