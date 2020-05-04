package generators

import "math/rand"

// MaleNameGenerator generates a first names
func MaleNameGenerator() string {
	return data[rand.Intn(len(data))]
}

var data = []string{
	"test",
}
