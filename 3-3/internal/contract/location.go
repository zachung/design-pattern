package contract

type Location struct {
	X uint
	Y uint
}

func (l *Location) Copy() *Location {
	return &Location{l.X, l.Y}
}

func (l *Location) Up() *Location {
	if l.Y != 0 {
		l.Y -= 1
	}
	return l
}

func (l *Location) Down() *Location {
	if l.Y+1 != Y {
		l.Y += 1
	}
	return l
}

func (l *Location) Left() *Location {
	if l.X != 0 {
		l.X -= 1
	}
	return l
}

func (l *Location) Right() *Location {
	if l.X+1 != X {
		l.X += 1
	}
	return l
}
