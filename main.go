package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/LarryC42/beelzebub/gen"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	profile := gen.ProfileGenerator()
	for k, v := range profile {
		fmt.Printf("%s: %s\n", k, v)
	}
}
