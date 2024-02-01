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
	Consume(chan<- QueueDto) error
}

type Queue struct {
	QueueConnection QueueConnection
}

func New(qt QueueType, config any) (q *Queue, err error){
  rf := reflect.TypeOf(config)

	switch qt {
  case QUEUE_TYPE_RABBITMQ:
		if rf.Name() != "RabbitMQ" {
      return nil, fmt.Errorf("Invalid config type")
    }

    conn, err := NewRabbitMQConnection(config.(RabbitMQ))
    if err != nil {
      return nil, err
    }
    q.QueueConnection = conn
	default:
		log.Fatal("Invalid Queue Type")
	}

	return q, nil
}

func (q *Queue) Publish(message []byte) error {
	return q.QueueConnection.Publish(message)
}

func (q *Queue) Consume(channelDto chan <- QueueDto) error {
  return q.QueueConnection.Consume(channelDto)
}
