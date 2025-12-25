package com.example.consumer;

import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.stereotype.Component;

@Component
public class RabbitMQConsumer {

    @RabbitListener(queues = "hello")
    public void receiveMessage(String message) {
        System.out.println(" [x] 接收到消息: " + message);
    }

    @RabbitListener(queues = "order")
    public void receiveOrderMessage(String message) {
        System.out.println(" [x] 订单接收到消息: " + message);
    }
}

