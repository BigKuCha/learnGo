package utils

import (
	"github.com/streadway/amqp"
	"log"
)

var (
	channel *amqp.Channel
)

const TopicExchange = "topicExchange"
const FanoutExchange = "fanoutExchange"

type myMsg struct {
	Body string
	Key  string
}

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
	//directMQ()
	//fanoutMQ()
}

//扇列模式，不需要路由，消息会发送到绑定此交换机上的所有队列
func fanoutMQ() {
	err := channel.ExchangeDeclare(FanoutExchange, "fanout", true, false, false, false, nil)
	if err != nil {
		log.Fatalln("交换机声明错误", err)
	}
	channel.QueueDeclare("fanoutQueue1", true, false, false, false, nil)
	channel.QueueDeclare("fanoutQueue2", true, false, false, false, nil)

	channel.QueueBind("fanoutQueue1", "", FanoutExchange, false, nil)
	channel.QueueBind("fanoutQueue2", "", FanoutExchange, false, nil)

	msg := amqp.Publishing{
		Body:[]byte("我是扇列消息"),
	}
	err = channel.Publish(FanoutExchange, "", false, false, msg)
	if err != nil {
		log.Fatalln("消息发送失败", err)
	}
}

//直连模式，不需要声明交换机，路由名称默认为队列名称，消息直接发送到相应队列
func directMQ() {
	directQueueName := "directQueue"
	//不需要声明交换机，默认使用rabbitmq的默认交换机
	channel.QueueDeclare(directQueueName, true, false, false, false, nil)
	msg := amqp.Publishing{
		Body:[]byte("我是消息体"),
	}
	err := channel.Publish("", directQueueName, false, false, msg) //自动创建`directQueueName`的交换机
	if err != nil {
		log.Fatalln("消息发送失败", err)
	}
}

//主题模式，消息会发送到该交换机上匹配路由的队列上
func topicMQ() {
	err := channel.ExchangeDeclare(TopicExchange, "topic", true, false, false, false, nil)
	if err != nil {
		log.Fatalln("交换机声明错误", err)
	}
	//声明四个队列，分别存储：北京、上海、新闻、天气相关信息
	channel.QueueDeclare("topicBeijing", true, false, false, false, nil)
	channel.QueueDeclare("topicShanghai", true, false, false, false, nil)
	channel.QueueDeclare("topicNews", true, false, false, false, nil)
	channel.QueueDeclare("topicWeather", true, false, false, false, nil)

	//分别绑定四个路由，分别收集 北京相关信息，上海相关信息，天气相关信息，新闻相关信息
	channel.QueueBind("topicBeijing", "bj.*", TopicExchange, false, nil)
	channel.QueueBind("topicShanghai", "shanghai.*", TopicExchange, false, nil)
	channel.QueueBind("topicNews", "*.news", TopicExchange, false, nil)
	channel.QueueBind("topicWeather", "*.weather", TopicExchange, false, nil)

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
