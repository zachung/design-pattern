package contract

type Video struct {
	Title       string
	Description string
	Length      int
	Channel     Subject
}

// Observer 是觀察者接口
type Observer interface {
	Update(event string, data *interface{})
	Observe(event string, f func(data *interface{}))
}

// Subject 是主題接口
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(event string, data *interface{})
}
