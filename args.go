package uu

import "os"

var Args []string
var ArgIndex int = -1

func init() {
	Args = os.Args[1:]
}

func Poll() string {
	if len(Args) == ArgIndex {
		panic("arg polling error")
	}
	ArgIndex += 1
	return Args[ArgIndex]
}

func MinArgs(minArgs int, errorMessage string) {
	if len(Args) < minArgs {
		panic(errorMessage)
	}
}

func Run(m map[string]func()) {
	if fn, ok := m[Poll()]; ok {
		fn()
	} else {
		panic("command not found")
	}
}
