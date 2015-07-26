package main

import "fmt"

type Error struct {
	err string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", APP_NAME, e.err)
}
