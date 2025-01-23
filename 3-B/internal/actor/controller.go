package actor

type Controller struct {
	commands   [][]int
	interactCh chan []int
}

func NewController() *Controller {
	return &Controller{
		commands:   make([][]int, 0),
		interactCh: make(chan []int),
	}
}

func (c *Controller) AddCommand(ints []int) {
	c.interactCh <- ints
}

func (c *Controller) PullCommand() []int {
	return <-c.interactCh
}
