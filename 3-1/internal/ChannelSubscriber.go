package internal

import "3-1/internal/contract"

type ChannelSubscriber struct {
	Name               string
	onNewVideoUploaded func(contract.Video)
}

func NewChannelSubscriber(name string, f func(video contract.Video)) contract.ChannelSubscriber {
	return &ChannelSubscriber{Name: name, onNewVideoUploaded: f}
}

func (c *ChannelSubscriber) Notify(video contract.Video) {
	c.onNewVideoUploaded(video)
}

func (c *ChannelSubscriber) GetName() string {
	return c.Name
}
