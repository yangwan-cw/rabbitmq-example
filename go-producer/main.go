package main

import (
	"log"
	"strconv"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// 加载配置
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 连接到RabbitMQ服务器
	conn, err := amqp.Dial(config.RabbitMQ.ConnectionString())
	if err != nil {
		log.Fatalf("无法连接到RabbitMQ: %v", err)
	}
	defer conn.Close()

	// 创建通道
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("无法打开通道: %v", err)
	}
	defer ch.Close()

	// 声明队列
 ch.QueueDeclare(
		config.RabbitMQ.Queue, // 队列名称
		false,                 // 持久化
		false,                 // 自动删除
		false,                 // 排他性
		false,                 // 无等待
		nil,                   // 参数
	)
	if err != nil {
		log.Fatalf("无法声明队列: %v", err)
	}

	// 发送消息
	// ch.Publish() 参数说明：
	// 参数1: exchange - 交换机名称
	//   - "" (空字符串): 默认 Exchange（Direct 类型）
	//   - "direct-exchange": Direct Exchange（直连交换机）
	//   - "topic-exchange": Topic Exchange（主题交换机）
	//   - "fanout-exchange": Fanout Exchange（扇出交换机）
	//   - "headers-exchange": Headers Exchange（头交换机）
	// 参数2: key - 路由键（routing key）
	// 参数3: mandatory - 强制（如果为true，消息无法路由时会返回错误）
	// 参数4: immediate - 立即（已废弃，应设为false）
	for i := 1; i <= 10; i++ {
		body := "Hello RabbitMQ! Message Number: " + strconv.Itoa(i)
		err = ch.Publish(
			"amq.fanout",     // 交换机：空字符串 = 默认 Exchange（Direct 类型）
			"", // 路由键：使用队列名称作为 routing key
			false,  // 强制：如果消息无法路由，不返回错误
			false,  // 立即：已废弃，设为 false
			amqp.Publishing{
				ContentType: "text/plain; charset=utf-8",
				Body:        []byte(body),
			})
		if err != nil {
			log.Fatalf("无法发布消息: %v", err)
		}
		log.Printf(" [x] 发送: %s", body)
		time.Sleep(1 * time.Second)
	}

	log.Println("所有消息已发送完成")
}

