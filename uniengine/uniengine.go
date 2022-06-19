package uniengine

import (
	"math/rand"
	"time"
)

type AsciiRange struct {
	start, end int
}

var (
	ArSymbol    = AsciiRange{32, 47}
	ArLowercase = AsciiRange{97, 122}
	ArUppercase = AsciiRange{65, 90}
	ArNumber    = AsciiRange{48, 57}
)

func aggregate(buffer *[]rune, r AsciiRange, size int, agg bool) []rune {
	if agg {
		*buffer = append(*buffer, RandomInRange(r, size)...)
	}
	return *buffer
}

func GenerateRandomKey(lower, upper, number, symbol bool, size int) string {
	args := sumTrueValues(lower, upper, number, symbol)
	last := size % args
	size_of_part := size / args
	buffer := make([]rune, 0)

	aggregate(&buffer, ArLowercase, size_of_part+last, lower)
	aggregate(&buffer, ArUppercase, size_of_part, upper)
	aggregate(&buffer, ArNumber, size_of_part, number)
	aggregate(&buffer, ArSymbol, size_of_part, symbol)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(buffer), func(i, j int) {
		buffer[i], buffer[j] = buffer[j], buffer[i]
	})

	return string(buffer)
}

func sumTrueValues(values ...bool) int {
	total := 0
	for _, v := range values {
		if v {
			total++
		}
	}
	return total
}

func RandomInRange(r AsciiRange, size int) []rune {
	buffer := make([]rune, size)
	for i := 0; i < size; i++ {
		data := rand.Intn(r.end-r.start) + r.start
		buffer[i] = rune(data)
	}

	return buffer
}
