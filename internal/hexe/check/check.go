package check

import "fmt"

func That(cond bool, format string, a ...any) {
	if !cond {
		panic(fmt.Errorf(format, a...))
	}
}

func Error(err error) {
	if err != nil {
		panic(err)
	}
}
