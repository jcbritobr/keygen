package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const (
	LowercaseBuffer = "abcdefghijklmnopqrstuvwxyz"
	UppercaseBuffer = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SymbolBuffer    = "()*&%$#@!;:/?.><,"
	NumberBuffer    = "0123456789"
)

var (
	keySizeFlag   = flag.Int("k", 14, "Use -k for the key size")
	lowercaseFlag = flag.Bool("l", true, "Use -l to insert lowercase characters in key")
	uppercaseFlag = flag.Bool("u", false, "Use -u to insert uppercase characters in key")
	symbolFlag    = flag.Bool("s", false, "Use -s to insert uppercase characters in key")
	numberFlag    = flag.Bool("n", false, "Use -n to insert uppercase characters in key")
)

func generateKey(lower, upper, symbol, number bool, size int) string {
	var buffer string

	if lower {
		buffer += LowercaseBuffer
	}
	if upper {
		buffer += UppercaseBuffer
	}
	if symbol {
		buffer += SymbolBuffer
	}
	if number {
		buffer += NumberBuffer + NumberBuffer // to make higher the number probability when all kinds of buffer is involved
	}

	rand.Seed(time.Now().Unix())
	runeBuffer := []rune(buffer)
	rand.Shuffle(len(runeBuffer), func(i, j int) {
		runeBuffer[i], runeBuffer[j] = runeBuffer[j], runeBuffer[i]
	})

	return string(runeBuffer[:size])
}
func main() {
	flag.Parse()
	key := generateKey(*lowercaseFlag, *uppercaseFlag, *symbolFlag, *numberFlag, *keySizeFlag)
	fmt.Println(key)
}
