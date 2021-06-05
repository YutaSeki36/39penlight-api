package mqtt

import (
	"errors"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"os"
)

const (
	mqttChannel = "penlight"
	resource    = "color"
)

type MQTTClient struct {
	client MQTT.Client
}

func NewMQTTClient() *MQTTClient {

	broker := os.Getenv("BROKER")
	mqttToken := os.Getenv("TOKEN")

	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetUsername(fmt.Sprintf("token:%s", mqttToken))
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return &MQTTClient{
		client: client,
	}
}

func (m *MQTTClient) Publish(massage string) bool {

	token := m.client.Publish(fmt.Sprintf("%s/%s", mqttChannel, resource), 0, false, massage)
	// TODO: Waitは無限に待ち状態となるのでWaitTimeoutに差し替える
	result := token.Wait()

	return result
}

func (m *MQTTClient) Disconnect() error {
	if m.client.IsConnected() {
		m.client.Disconnect(250)
		fmt.Println("client disconnected")
		return nil
	}
	return errors.New("client has no connection")
}
