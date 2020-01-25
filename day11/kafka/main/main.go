/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/23 6:48 下午
 *  Notes       :	Kafka example
./zkServer.sh start
./kafka-console-consumer.sh --topic nginx_log --bootstrap-server localhost:9092
./kafka-server-start.sh /Users/nieaowei/Desktop/Environment/kafka_2.12-2.4.0/config/server.properties
 *******************************************************/

package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	//Kafka config
	saramaCfg := sarama.NewConfig()
	saramaCfg.Producer.RequiredAcks = sarama.WaitForAll
	saramaCfg.Producer.Partitioner = sarama.NewRandomPartitioner
	saramaCfg.Producer.Return.Successes = true
	// Client
	saramaCli, err := sarama.NewSyncProducer([]string{"localhost:9092"}, saramaCfg)
	if err != nil {
		fmt.Println("produer close , err:", err)
		return
	}
	defer saramaCli.Close()

	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx_log"
	msg.Value = sarama.StringEncoder("this is test.")

	pid, offet, err := saramaCli.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed,", err)
		return
	}
	fmt.Println("pid:", pid, "offet:", offet)
	return
}
