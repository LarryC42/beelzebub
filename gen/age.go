package gen

import (
	"fmt"
	"math"
	"math/rand"
)

// RandomAgeGenerator generates an age for a child (0-96)
func RandomAgeGenerator() string {
	return fmt.Sprintf("%d", rand.Intn(97))
}

// ChildAgeGenerator generates an age for a child (0-17)
func ChildAgeGenerator() string {
	return fmt.Sprintf("%d", rand.Intn(18))
}

// WorkingAdultAgeGenerator generates an age for a working adult+ (18-67)
func WorkingAdultAgeGenerator() string {
	return fmt.Sprintf("%d", rand.Intn(68-18)+18)
}

// RetiredAgenGenerator generates an age for a retired person (60-96)
func RetiredAgenGenerator() string {
	return fmt.Sprintf("%d", rand.Intn(97-60)+60)
}

// RaceGenerator generates a race for a person
func RaceGenerator() string {
	return raceData[rand.Intn(len(raceData))]
}

var raceData = []string{
	"Native American", "Asian", "Black", "Pacific Islander", "Hispanic", "White",
}

func isValidSsn(ssn string) bool {
	if len(ssn) != 9 {
		return false
	}
	if ssn[0] == '9' {
		return false
	}
	if ssn[:3] == "666" {
		return false
	}
	if ssn[:3] == "000" {
		return false
	}
	if ssn[3:5] == "00" {
		return false
	}
	if ssn[5:] == "000" {
		return false
	}
	return true
}

// SsnGenerator generates a Social Security Number
func SsnGenerator() string {
	for true {
		ssn := fmt.Sprintf("%9.9d", rand.Intn(1000000000))
		if isValidSsn(ssn) {
			return fmt.Sprintf("%s-%s-%s", ssn[:3], ssn[3:5], ssn[5:])
		}
	}
	return ""
}

// EyeColorGenerator generates a race for a person
func EyeColorGenerator() string {
	return eyeColorData[rand.Intn(len(eyeColorData))]
}

var eyeColorData = []string{
	"brown", "brown", "brown", "brown", "brown", "brown", "brown", "blue", "gray", "hazel", "green", "amber",
}

// MaleHeightInches generates male height
func MaleHeightInches() int {
	return (int)(math.Max(math.Min(rand.NormFloat64()*(11.0)+70.0, 78), 56)) // Male height
}

// FemaleHeightInches generates female height
func FemaleHeightInches() int {
	return (int)(math.Max(math.Min(rand.NormFloat64()*(11.0)+64.0, 78), 56)) // Female height
}

// InchesToHeight converts inches to x' y"
func InchesToHeight(height int) string {
	var ft = (int)(height / 12)
	var in = (int)(height % 12)
	return fmt.Sprintf("%d' %d\"", ft, in)
}

// MaleWeightLbs generates male weight
func MaleWeightLbs(height int) string {
	var lbs = (int)(math.Max(math.Min(rand.NormFloat64()*(10.0)+((float64)(height)*2.8), 300.0), 130.0)) // Male weight
	return fmt.Sprintf("%d", lbs)
}

// FemaleWeightLbs generates female weight
func FemaleWeightLbs(height int) string {
	var lbs = (int)(math.Max(math.Min(rand.NormFloat64()*(10.0)+((float64)(height)*2.66), 280.0), 110.0)) // Female weight
	return fmt.Sprintf("%d", lbs)
}
