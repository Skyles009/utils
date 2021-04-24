package subscriber

import (
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

func SubscribeToQueue(subject string, queue string, options []stan.SubscriptionOption, handler stan.MsgHandler) stan.Subscription {
	sub, _ := StanConn.QueueSubscribe(subject, queue, handler, options...)

	return sub
}
