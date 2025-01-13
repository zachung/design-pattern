package internal

type Property struct {
	value    int
	observer []func(*int)
}

func NewProperty(value int) *Property {
	return &Property{
		value:    value,
		observer: make([]func(*int), 0),
	}
}

func (p *Property) Get() int {
	return p.value
}

func (p *Property) Sub(v int) {
	p.value -= v
	for _, f := range p.observer {
		f(&p.value)
	}
}

func (p *Property) Add(v int) {
	p.value += v
}

func (p *Property) AddObserver(observer func(*int)) {
	p.observer = append(p.observer, observer)
}
