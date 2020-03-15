package rabbitmqworker

import (
	"Edwardz43/tgbot/config"
	"Edwardz43/tgbot/err"
	"Edwardz43/tgbot/message/from"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

var failOnError = err.FailOnError

// Worker is the worker uses rabbitmq
type Worker struct {
	channel   *amqp.Channel
	queueName string
	Result    *from.Result
	//Job       *worker.Job

}

// connect creates a rabbitmq client connection
func (r *Worker) connect() bool {

	conn, err := amqp.Dial(config.GetRabbitDNS())

	if err != nil {
		log.Printf("Failed to connect to RabbitMQ : %s", err)
		return false
	}

	//defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Failed to open a channel : %s", err)
		return false
	}

	q, err := ch.QueueDeclare(
		"tgbot_message", // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)

	if err != nil {
		log.Printf("Failed to declare a queue : %s", err)
		return false
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	if err != nil {
		log.Printf("Failed to set QoS : %s", err)
		return false
	}

	r.channel = ch

	r.queueName = q.Name

	return true
}

// Do executes job
func (r *Worker) Do(job func(args ...interface{}) error) {
	ok := r.connect()

	if !ok {
		log.Panicf("Failed to connect to rabbitmq channel")
		return
	}

	msgs, err := r.channel.Consume(
		r.queueName, // queue
		"",          // consumer
		false,       // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)

	failOnError(err, "Failed to register a consumer")

	defer r.channel.Close()

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Worker Received a message: %s", d.Body)
			json.Unmarshal(d.Body, &r.Result)
			if r.Result == nil {
				log.Panicln("json unmarshal failed")
			}
			err := job(r.Result)
			//TODO
			failOnError(err, "Failed : error from job")
			log.Println("Worker Done")
			d.Ack(false)
		}
	}()

	<-forever
}
