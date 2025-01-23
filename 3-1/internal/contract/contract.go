package contract

type Channel interface {
	Subscribe(subscriber ChannelSubscriber)
	UnSubscribe(subscriber ChannelSubscriber)
	Upload(video Video)
}

type ChannelSubscriber interface {
	Notify(video Video)
	GetName() string
}

type Video struct {
	Title       string
	Description string
	Length      int
	Channel     Channel
}
