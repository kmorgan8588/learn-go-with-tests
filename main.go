package main

import (
	"learn-go-with-tests/mocking"
	"os"
)

func main() {
	mocking.Countdown(os.Stdout)

}
