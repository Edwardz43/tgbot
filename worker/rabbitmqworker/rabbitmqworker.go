package rabbitmqworker

import (
	"Edwardz43/tgbot/config"
	"Edwardz43/tgbot/err"
	"Edwardz43/tgbot/log"
	"Edwardz43/tgbot/message/from"
	"Edwardz43/tgbot/worker"
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

var failOnError = err.FailOnError

// GetInstance returns a instance of rabbitmq worker
func GetInstance(l log.Logger) worker.Worker {
	return &Worker{
		logger: l,
	}
}

// Worker is the worker uses rabbitmq
type Worker struct {
	channel   *amqp.Channel
	queueName string
	Result    *from.Result
	logger    log.Logger
}

// connect creates a rabbitmq client connection
func (r *Worker) connect() bool {

	conn, err := amqp.Dial(config.GetRabbitDNS())

	if err != nil {
		r.logger.ERROR(fmt.Sprintf("Failed to connect to RabbitMQ : %s", err))
		return false
	}

	//defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		r.logger.ERROR(fmt.Sprintf("Failed to open a channel : %s", err))
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
		r.logger.ERROR(fmt.Sprintf("Failed to declare a queue : %s", err))
		return false
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	if err != nil {
		r.logger.ERROR(fmt.Sprintf("Failed to set QoS : %s", err))
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
		r.logger.ERROR(fmt.Sprintln("Failed to connect to rabbitmq channel"))
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
			r.logger.INFO(fmt.Sprintf("Worker Received a message: %s", d.Body))
			json.Unmarshal(d.Body, &r.Result)
			if r.Result == nil {
				r.logger.PANIC("json unmarshal failed")
			}
			err := job(r.Result)
			failOnError(err, "Failed : error from job")
			r.logger.INFO("Work Done")
			d.Ack(false)
		}
	}()

	<-forever
}
