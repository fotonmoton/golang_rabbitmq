package notifications

import (
	"context"
	"docker/warehouse"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var customerNotificationsQueue = "customerNotifications"

type RabbitMQChannel struct {
	conn *amqp.Connection
}

func NewRabbitMQChannel() *RabbitMQChannel {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")

	failOnError(err, "Failed to connect to RabbitMQ")

	return &RabbitMQChannel{conn}
}

func (r *RabbitMQChannel) NotifyCustomers(notifications []warehouse.CustomerNotification) {

	ch, err := r.conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		customerNotificationsQueue, // name
		false,                      // durable
		false,                      // delete when unused
		false,                      // exclusive
		false,                      // no-wait
		nil,                        // arguments
	)

	failOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, notification := range notifications {
		body, _ := json.Marshal(notification)
		err = ch.PublishWithContext(ctx,
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        body,
			})
		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s\n", body)
	}
}

func (r *RabbitMQChannel) ProcessNotifications() {

	ch, err := r.conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		customerNotificationsQueue, // name
		false,                      // durable
		false,                      // delete when unused
		false,                      // exclusive
		false,                      // no-wait
		nil,                        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for message := range msgs {
			notification := warehouse.CustomerNotification{}
			json.Unmarshal(message.Body, &notification)
			sendEmail(notification)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func sendEmail(notification warehouse.CustomerNotification) {
	log.Printf("Received a message: %+v", notification)
}
