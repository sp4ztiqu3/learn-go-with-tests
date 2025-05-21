package main

import (
	"os"

	"github.com/sp4ztiqu3/learn-go-with-tests/mocking"
)

func main() {
	sleeper := &mocking.DefaultSleeper{}
	mocking.Countdown(os.Stdout, sleeper)
}
