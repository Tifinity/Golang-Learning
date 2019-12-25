package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生产者一直每隔一秒像通道发送一个数据
func producer(header string, channel chan <-string) {
	for {
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31())
		time.Sleep(time.Second)
	}
}

// 消费者从通道中阻塞地取出数据
func customer(channel <-chan string) {
    for {
        message := <-channel
        fmt.Println(message)
    }
}

func main() {
	channel := make(chan string)
	go producer("cat", channel)
	go producer("dog", channel)
	customer(channel)
}