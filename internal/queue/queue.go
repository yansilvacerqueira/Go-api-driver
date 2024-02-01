package queue

import "log"
import "reflect"
import "fmt"

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

func New(qt QueueType, config any) (q *Queue, err error){
  rf := reflect.TypeOf(config)

	switch qt {
  case QUEUE_TYPE_RABBITMQ:
		if rf.Name() != "RabbitMQ" {
      return nil, fmt.Errorf("Invalid config type")
    }

    q.QueueConnection = NewRabbitMQConnection(config)
	default:
		log.Fatal("Invalid Queue Type")
	}

	return 
}

func (q *Queue) Publish(message []byte) error {
	return q.QueueConnection.Publish(message)
}

func (q *Queue) Consume() error {
  return q.QueueConnection.Consume()
}
