package queue

import "time"
import amqp "github.com/rabbitmq/amqp091-go"
import "context"

type RabbitMQ struct {
  URL string
  TopicName string
  Timeout time.Time
}

type RabbitMQConnection struct { 
  config RabbitMQ
  conn *amqp.Connection
}

func NewRabbitMQConnection(config RabbitMQ) (rc *RabbitMQConnection, err error) {
  rc.config = config
  
  rc.conn, err = amqp.Dial(config.URL)

  return rc, err
}

func (rc *RabbitMQConnection) Publish(message []byte) error {
  ch, err := rc.conn.Channel()
  if err != nil {
    return err
  }
  
  messagePublish := amqp.Publishing{
    ContentType: "text/plain",
    Body:        message,
    DeliveryMode: amqp.Persistent,
    Timestamp: time.Now(),
  }
  
  ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
  
  defer cancel()

  return ch.PublishWithContext(ctx, "", rc.config.TopicName, false, false, messagePublish)
}

func (rc *RabbitMQConnection) Consume(channelDto chan <- QueueDto) error {
  ch, err := rc.conn.Channel()
  if err != nil {
    return err  
  }
  
  qeue, err := ch.QueueDeclare(rc.config.TopicName, false, false, false, false, nil)
  if err != nil {
    return err
  }
  
  messages, err := ch.Consume(qeue.Name, "", true, false, false, false, nil)
  if err != nil {
    return err  
  }

  for message := range messages {
    
    dto := QueueDto{
      FileName: message.Exchange,
      Path: message.RoutingKey,
    } 
    dto.Unmarshal(message.Body)
    channelDto <- dto
  }

  return nil
}
