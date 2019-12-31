package main

import (
	"./pubsub"
	"time"
	"strings"
	"fmt"
	"syscall"
	"os/signal"
	"os"
)

func main() {
	p := pubsub.NewPublisher(100 * time.Millisecond, 10)

	golang := p.SubscribeTopic(func (v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "golang:") {
				return true
			}
		}
		return false
	})

	docker := p.SubscribeTopic(func (v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "docker:") {
				return true
			}
		}
		return false
	})

	go p.Publish("hi")
	go p.Publish("golang: https://golang.org")
	go p.Publish("docker: https://www.docker.com/")

	go func() {
		fmt.Println("golang topic:", <-golang)
	}()

	go func() {
		fmt.Println("docker topic:", <-docker)
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)
	fmt.Printf("quit (%v)\n", <-sig)
}