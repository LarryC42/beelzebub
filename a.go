package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 20; i++ {
		fmt.Println(rand.NormFloat64() * .5)
	}
}
