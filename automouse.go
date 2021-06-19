package main

import (
	"github.com/go-vgo/robotgo"
  "math/rand"
  "time"
)

func main() {
  for range time.Tick(time.Minute * 1) {
    go doMouseStuffs()
  }
}

func doMouseStuffs() {
  robotgo.MoveSmooth(rand.Intn(1000), rand.Intn(500))
}