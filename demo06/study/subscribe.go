package study

import (
	"fmt"
	"sync"
	"time"
)

type (
	subscriber chan interface{}         // 订阅者为一个管道
	topicFunc  func(v interface{}) bool // 主题为一个过滤器
)

// 发布者对象
type Publisher struct {
	mutex      sync.Mutex
	buffer     int
	timeout    time.Duration
	subscriber map[subscriber]topicFunc
}

// 构建一个发布者对象，设置发布超时时间和缓存队列
func NewPublisher(timeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:     buffer,
		timeout:    timeout,
		subscriber: make(map[subscriber]topicFunc),
	}
}

// 添加一个新的订阅者，订阅全部主题
func (p *Publisher) Subscriber() chan interface{} {
	return p.SubscriberTopic(nil)
}

// 添加一个新的订阅者，定义特定主题
func (p *Publisher) SubscriberTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.mutex.Lock()
	p.subscriber[ch] = topic
	p.mutex.Unlock()
	return ch
}

// 退出订阅
func (p *Publisher) Evict(ch chan interface{}) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	delete(p.subscriber, ch)
}

// 发送
func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {

	defer wg.Done()

	if topic != nil && !topic(v) {
		return
	}

	select {
	case sub <- v:
	case <-time.After(p.timeout):
		fmt.Println("超时")
	}
}

// 发布主题
func (p *Publisher) Publish(v interface{}) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	var wg sync.WaitGroup

	for sub, topic := range p.subscriber {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

// 关闭
func (p *Publisher) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for sub := range p.subscriber {
		close(sub)
	}

}

func TestPublish() {
	p := NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	allTopic := p.Subscriber()

	golang := p.SubscriberTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return s == "golang"
		}
		return false
	})

	p.Publish("hello world!")
	p.Publish("golang")

	go func() {
		for msg := range allTopic {
			fmt.Println("allTopic:", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang:", msg)
		}
	}()

	time.Sleep(3 * time.Second)
}
