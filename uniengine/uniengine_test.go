package uniengine

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

type runeKind uint8

const (
	rkNumber runeKind = iota
	rkUppercase
	rkLowercase
	rkSymbol
)

func assertAlmostOneRuneKind(t *testing.T, kind runeKind, buffer string) bool {
	t.Helper()
	for _, r := range buffer {
		switch kind {
		case rkNumber:
			if unicode.IsDigit(r) {
				return true
			}
		case rkUppercase:
			if unicode.IsUpper(r) {
				return true
			}
		case rkLowercase:
			if unicode.IsLower(r) {
				return true
			}
		case rkSymbol:
			if unicode.IsPunct(r) {
				return true
			}
		}
	}

	return false
}

func assertRuneKind(t *testing.T, kind runeKind, buffer string) bool {
	t.Helper()
	result := false

	for _, r := range buffer {
		switch kind {
		case rkNumber:
			result = unicode.IsDigit(r)
		case rkUppercase:
			result = unicode.IsUpper(r)
		case rkLowercase:
			result = unicode.IsLower(r)
		case rkSymbol:
			result = unicode.IsPunct(r)
		}
	}

	return result
}

func TestRandomInRangeIsCharacterKindAndLength(t *testing.T) {
	// Given I'm a component,
	// when I generate 100 random runes
	// and runes are choosed to be lowercase letters
	result := string(RandomInRange(ArLowercase, 100))
	assert.Equal(t, 100, len(result))
	boolResult := assertRuneKind(t, rkLowercase, result)
	assert.True(t, boolResult)
	// and I generate 10 runes and runes are choosed to be uppercase letters
	result = string(RandomInRange(ArUppercase, 10))
	assert.Equal(t, 10, len(result))
	boolResult = assertRuneKind(t, rkUppercase, result)
	assert.True(t, boolResult)
	// and I generate 1000 runes and runes are choosed to be numbers
	result = string(RandomInRange(ArNumber, 1000))
	assert.Equal(t, 1000, len(result))
	boolResult = assertRuneKind(t, rkNumber, result)
	assert.True(t, boolResult)
	// and I generate 300 runes and runes are choosed to be symbols
	result = string(RandomInRange(ArSymbol, 300))
	assert.Equal(t, 300, len(result))
	boolResult = assertRuneKind(t, rkSymbol, result)
	assert.True(t, boolResult)
	// then must not have any kind of errors
}

func TestSumTrueValue(t *testing.T) {
	// Given I'm a component
	// When I pass 5 boolean values to sumTrueValue(...), 2 falses and 3 trues
	result := sumTrueValues(false, false, true, true, true)
	// Then must return 3 as result
	expected := 3
	assert.Equal(t, expected, result)
}

func TestGenerateKeyFunction(t *testing.T) {
	// Given I'm a component
	// When I generate a 16 size random key with (true, false, false, false) params
	key := GenerateRandomKey(true, false, false, false, 16)
	// Then key must have 16 len size
	expected := 16
	assert.Equal(t, expected, len(key))
	// And all runes must be lowercase
	assertRuneKind(t, rkLowercase, key)
	// same below with upper and size 100
	key = GenerateRandomKey(false, true, false, false, 100)
	expected = 100
	assert.Equal(t, expected, len(key))
	assertRuneKind(t, rkUppercase, key)
	// same below with number and size 120
	key = GenerateRandomKey(false, false, true, false, 120)
	expected = 120
	assert.Equal(t, expected, len(key))
	assertRuneKind(t, rkNumber, key)
	// same below with symbol and size 8
	key = GenerateRandomKey(false, false, false, true, 8)
	expected = 8
	assert.Equal(t, expected, len(key))
	assertRuneKind(t, rkSymbol, key)
	// Given I'm a component
	// When I generate a key with symbols and lowercase letters, with size 8
	key = GenerateRandomKey(true, false, false, true, 8)
	expected = 8

	// Then the key must have symbols and lowercase letters, and size 8
	assert.Equal(t, expected, len(key))
	boolResult := assertAlmostOneRuneKind(t, rkLowercase, key)
	assert.True(t, boolResult)
	boolResult = assertAlmostOneRuneKind(t, rkSymbol, key)
	assert.True(t, boolResult)
	// When I generate a key with uppercase letters and numbers, with size 16
	key = GenerateRandomKey(false, true, true, false, 16)
	expected = 16
	assert.Equal(t, expected, len(key))
	// Then the key must have uppercase letters and numbers, and size 16
	boolResult = assertAlmostOneRuneKind(t, rkNumber, key)
	assert.True(t, boolResult)
	boolResult = assertAlmostOneRuneKind(t, rkUppercase, key)
	assert.True(t, boolResult)
}
