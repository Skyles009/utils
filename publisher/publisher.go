package publisher

import (
	"encoding/json"
	"log"

	"github.com/Skyles009/utils/randomString"
	"github.com/nats-io/stan.go"
)

var StanConn stan.Conn
var err error

func InitialiseNatsConn(url string, serverID string, charset string) {
	clientID := randomString.String(7, charset)
	StanConn, err = stan.Connect(serverID, clientID, stan.NatsURL(url))
	if err != nil {
		log.Fatal(err.Error())
	}

}

func PublishToQueue(subject string, m interface{}) {
	message, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = StanConn.Publish(subject, message)
	if err != nil {
		log.Fatal(err.Error())
	}

}
