package main

import (
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
)

// 示例：如何使用不同类型的 Exchange

func exampleDirectExchange(ch *amqp.Channel) {
	// 1. Direct Exchange（直连交换机）- 默认类型
	// 路由规则：完全匹配 routing key
	err := ch.ExchangeDeclare(
		"direct-exchange", // Exchange 名称
		"direct",          // Exchange 类型
		false,             // 持久化
		false,             // 自动删除
		false,             // 内部（不接收外部发布）
		false,             // 无等待
		nil,               // 参数
	)
	if err != nil {
		log.Fatalf("声明 Direct Exchange 失败: %v", err)
	}

	// 发布消息到 Direct Exchange
	err = ch.Publish(
		"direct-exchange", // Exchange 名称
		"order.create",    // routing key（必须完全匹配）
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Direct Exchange 消息"),
		})
	if err != nil {
		log.Fatalf("发布消息失败: %v", err)
	}
}

func exampleTopicExchange(ch *amqp.Channel) {
	// 2. Topic Exchange（主题交换机）
	// 路由规则：支持通配符匹配 routing key
	// * 匹配一个单词， # 匹配零个或多个单词
	err := ch.ExchangeDeclare(
		"topic-exchange", // Exchange 名称
		"topic",         // Exchange 类型
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("声明 Topic Exchange 失败: %v", err)
	}

	// 发布消息到 Topic Exchange
	err = ch.Publish(
		"topic-exchange",      // Exchange 名称
		"order.create.vip",    // routing key（支持通配符匹配）
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Topic Exchange 消息"),
		})
	if err != nil {
		log.Fatalf("发布消息失败: %v", err)
	}
}

func exampleFanoutExchange(ch *amqp.Channel) {
	// 3. Fanout Exchange（扇出交换机）
	// 路由规则：忽略 routing key，广播到所有绑定的队列
	err := ch.ExchangeDeclare(
		"fanout-exchange", // Exchange 名称
		"fanout",          // Exchange 类型
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("声明 Fanout Exchange 失败: %v", err)
	}

	// 发布消息到 Fanout Exchange
	// routing key 会被忽略，消息会发送到所有绑定的队列
	err = ch.Publish(
		"fanout-exchange", // Exchange 名称
		"",                // routing key（会被忽略）
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Fanout Exchange 消息（广播）"),
		})
	if err != nil {
		log.Fatalf("发布消息失败: %v", err)
	}
}

func exampleHeadersExchange(ch *amqp.Channel) {
	// 4. Headers Exchange（头交换机）
	// 路由规则：根据消息头（headers）匹配，忽略 routing key
	err := ch.ExchangeDeclare(
		"headers-exchange", // Exchange 名称
		"headers",          // Exchange 类型
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("声明 Headers Exchange 失败: %v", err)
	}

	// 发布消息到 Headers Exchange
	// routing key 会被忽略，使用 headers 进行路由
	err = ch.Publish(
		"headers-exchange", // Exchange 名称
		"",                 // routing key（会被忽略）
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Headers: amqp.Table{
				"type":    "order",
				"priority": "high",
			},
			Body: []byte("Headers Exchange 消息"),
		})
	if err != nil {
		log.Fatalf("发布消息失败: %v", err)
	}
}

// 默认 Exchange（空字符串）
func exampleDefaultExchange(ch *amqp.Channel, queueName string) {
	// 使用默认 Exchange（空字符串）
	// 类型：Direct Exchange
	// 路由规则：routing key 必须等于队列名称
	err := ch.Publish(
		"",         // 空字符串 = 默认 Exchange
		queueName,  // routing key = 队列名称
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("默认 Exchange 消息"),
		})
	if err != nil {
		log.Fatalf("发布消息失败: %v", err)
	}
}

