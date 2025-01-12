package state

import "3B/internal/contract"

func GetState(name string) contract.State {
	switch name {
	case "正常":
		return NewNormalState()
	case "石化":
		return NewPetrochemical()
	}
	panic("state " + name + " not found")
}
