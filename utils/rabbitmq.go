package utils

import (
	"github.com/streadway/amqp"
	"log"
)

var (
	conn    *amqp.Connection
	channel *amqp.Channel
)

const TopicExchange = "topicExchange"

func init() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatalln("消息队列连接失败！", err)
	}
	channel, err = conn.Channel()
	if err != nil {
		log.Fatalln("通道开启失败", err)
	}
}

func TestRabbitMQ() {
	topicMQ()
}

type myMsg struct {
	Body string
	Key  string
}

//topic模式消息队列
func topicMQ() {
	err := channel.ExchangeDeclare(TopicExchange, "topic", true, false, false, false, nil)
	if err != nil {
		log.Fatalln("交换机声明错误", err)
	}
	//声明四个队列，分别存储：北京、上海、新闻、天气相关信息
	channel.QueueDeclare("beijing", true, false, false, false, nil)
	channel.QueueDeclare("shanghai", true, false, false, false, nil)
	channel.QueueDeclare("news", true, false, false, false, nil)
	channel.QueueDeclare("weather", true, false, false, false, nil)

	//分别绑定四个路由，分别收集 北京相关信息，上海相关信息，天气相关信息，新闻相关信息
	channel.QueueBind("beijing", "bj.*", TopicExchange, false, nil)
	channel.QueueBind("shanghai", "shanghai.*", TopicExchange, false, nil)
	channel.QueueBind("news", "*.news", TopicExchange, false, nil)
	channel.QueueBind("weather", "*.weather", TopicExchange, false, nil)

	//发送N条主题（城市天气、城市新闻）消息
	msgs := []myMsg{
		{"北京天气：有雾霾", "bj.weather"}, //消息会分发到'北京'和'天气'两个主题队列上
		{"北京新闻：中关村改造", "bj.news"},
		{"上海天气：晴天", "shanghai.weather"},
		{"上海新闻：浦东机场坠落一架飞机", "shanghai.news"},
	}
	for _, v := range msgs {
		msg := amqp.Publishing{
			Body: []byte(v.Body),
		}
		err := channel.Publish(TopicExchange, v.Key, false, false, msg)
		if err != nil {
			log.Fatalln("消息发送失败", err)
		}
	}
}
