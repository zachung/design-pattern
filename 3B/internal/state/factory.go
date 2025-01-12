package state

import "3B/internal/contract"

func GetState(name string) contract.State {
	switch name {
	case "正常":
		return NewNormalState()
	}
	panic("state " + name + " not found")
}
