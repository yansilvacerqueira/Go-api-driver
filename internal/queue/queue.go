package queue

import "log"

const (
	QUEUE_TYPE_RABBITMQ QueueType = iota
)

type QueueType int

type QueueConnection interface {
	Publish([]byte) error
	Consume() error
}

type Queue struct {
	config          any
	QueueConnection QueueConnection
}

func New(qt QueueType, config any) *Queue {
	q := &Queue{config: config}
	switch qt {
	case QUEUE_TYPE_RABBITMQ:
		q.QueueConnection = NewRabbitMQConnection(config)
	default:
		log.Fatal("Invalid Queue Type")
	}
	return q
}

func (q *Queue) Publish(message []byte) error {
	return q.QueueConnection.Publish(message)
}

func (q *Queue) Consume() error {
  return q.QueueConnection.Consume()
}
