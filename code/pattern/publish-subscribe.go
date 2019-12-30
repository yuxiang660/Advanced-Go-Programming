package main

import (
	"fmt"
	"sync"
	"time"
	"strings"
)

type (
	subscriberChannel chan string
	topicFilter func(contents string) bool
)

type publisher struct {
	sync.RWMutex
	bufferSize int
	timeout time.Duration
	subscribers map[subscriberChannel]topicFilter
}

func createPublisher(timeout time.Duration, bufferSize int) *publisher {
	return &publisher{
		bufferSize: bufferSize,
		timeout: timeout,
		subscribers: make(map[subscriberChannel]topicFilter),
	}
}

func (p *publisher) subscribeAll() subscriberChannel {
	return p.subscribeTopicFilter(nil)
}

func (p *publisher) subscribeTopicFilter(filter topicFilter) subscriberChannel {
	ch := make(chan string, p.bufferSize)
	p.Lock()
	defer p.Unlock()

	p.subscribers[ch] = filter
	return ch
}

func (p *publisher) evict(subscriber subscriberChannel) {
	p.Lock()
	defer p.Unlock()

	delete(p.subscribers, subscriber)
	close(subscriber)
}

func (p *publisher) sendTopic(subscriber subscriberChannel, filter topicFilter, contents string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	if filter != nil && !filter(contents) {
		return
	}

	select {
	case subscriber <- contents:
	case <-time.After(p.timeout):
	}
}

func (p *publisher) close() {
	p.Lock()
	defer p.Unlock()

	for subscriber := range p.subscribers {
		delete(p.subscribers, subscriber)
		close(subscriber)
	}
}

func (p *publisher) publish(contents string) {
	p.RLock()
	defer p.RUnlock()

	var waitGroup sync.WaitGroup
	for subscriber, topicFilter := range p.subscribers {
		waitGroup.Add(1)
		go p.sendTopic(subscriber, topicFilter, contents, &waitGroup)
	}
	waitGroup.Wait()
}

func main() {
	fmt.Println("Publish and Subscribe Pattern")

	p := createPublisher(100 * time.Millisecond, 10)
	defer p.close()

	subscriberForAll := p.subscribeAll()
	subscriberForGolang := p.subscribeTopicFilter(func(contents string) bool {
		return strings.Contains(contents, "golang")
	})

	p.publish("hello, world")
	p.publish("hello, golang")

	go func() {
		for msg := range subscriberForAll {
			fmt.Println("all:", msg)
		}
	}()

	go func() {
		for msg := range subscriberForGolang {
			fmt.Println("golang:", msg)
		}
	}()

	time.Sleep(3 * time.Second)
}
