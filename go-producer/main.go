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
	q, err := ch.QueueDeclare(
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
	for i := 1; i <= 10; i++ {
		body := "Hello RabbitMQ! 消息编号: " + strconv.Itoa(i)
		err = ch.Publish(
			"",     // 交换机
			q.Name, // 路由键
			false,  // 强制
			false,  // 立即
			amqp.Publishing{
				ContentType: "text/plain",
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

