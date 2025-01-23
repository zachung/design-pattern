package internal

import (
	"3-1/internal/contract"
	"log"
	"sync"
)

type Channel struct {
	Name      string
	observers sync.Map
}

func NewSubject(name string) contract.Subject {
	return &Channel{Name: name, observers: sync.Map{}}
}

func (c *Channel) RegisterObserver(observer contract.Observer) {
	log.Printf("%s 訂閱了 %s。\n", observer.(*ChannelSubscriber).Name, c.Name)
	c.observers.Store(observer, struct{}{})
}

func (c *Channel) RemoveObserver(observer contract.Observer) {
	log.Printf("%s 解除訂閱了 %s。", observer.(*ChannelSubscriber).Name, c.Name)
	c.observers.Delete(observer)
}

func (c *Channel) NotifyObservers(event string, data *interface{}) {
	video := (*data).(*contract.Video)
	log.Printf("頻道 %s 上架了一則新影片 \"%s\"。", c.Name, video.Title)
	(video).Channel = c
	c.observers.Range(func(key, value interface{}) bool {
		observer := key.(contract.Observer)
		observer.Update(event, data)
		return true
	})
}
