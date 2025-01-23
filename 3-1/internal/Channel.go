package internal

import (
	"3-1/internal/contract"
	"log"
)

type Channel struct {
	Name        string
	subscribers []contract.ChannelSubscriber
}

func NewChannel(name string) *Channel {
	return &Channel{name, make([]contract.ChannelSubscriber, 0)}
}

func (c *Channel) Subscribe(subscriber contract.ChannelSubscriber) {
	log.Printf("%s 訂閱了 %s。\n", subscriber.GetName(), c.Name)
	c.subscribers = append(c.subscribers, subscriber)
}

func (c *Channel) UnSubscribe(subscriber contract.ChannelSubscriber) {
	log.Printf("%s 解除訂閱了 %s。", subscriber.GetName(), c.Name)
	for i, s := range c.subscribers {
		if s.GetName() == subscriber.GetName() {
			c.subscribers = append(c.subscribers[:i], c.subscribers[i+1:]...)
			break
		}
	}
}

func (c *Channel) Upload(video contract.Video) {
	log.Printf("頻道 %s 上架了一則新影片 \"%s\"。", c.Name, video.Title)
	video.Channel = c
	for _, subscriber := range c.subscribers {
		subscriber.Notify(video)
	}
}
