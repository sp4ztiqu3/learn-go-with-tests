package main

import (
	"os"

	"github.com/sp4ztiqu3/learn-go-with-tests/mocking"
)

func main() {
	mocking.Countdown(os.Stdout)
}
