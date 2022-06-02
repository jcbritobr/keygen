package main

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestGenerateKeyMustGenerateLowercaseString(t *testing.T) {
	lowerStr := generateKey(true, false, false, false, 5)
	has_lowercase := false
	for _, c := range lowerStr {
		if unicode.IsLower(c) {
			has_lowercase = true
		}
	}

	assert.True(t, has_lowercase)
}

func TestGenerateKeyMustGenerateUppercaseString(t *testing.T) {
	upperStr := generateKey(false, true, false, false, 5)
	has_uppercase := false
	for _, c := range upperStr {
		if unicode.IsUpper(c) {
			has_uppercase = true
		}
	}

	assert.True(t, has_uppercase)
}

func TestGenerateKeyMustGenerateNumberString(t *testing.T) {
	numberStr := generateKey(false, false, false, true, 5)
	has_number := false
	for _, c := range numberStr {
		if unicode.IsDigit(c) {
			has_number = true
		}
	}

	assert.True(t, has_number)
}

func TestGenerateKeyMustGenerateSymbolString(t *testing.T) {
	symbolStr := generateKey(false, false, true, false, 5)
	has_symbol := false
	for _, c := range symbolStr {
		if unicode.IsPunct(c) {
			has_symbol = true
		}
	}

	assert.True(t, has_symbol)
}

func TestGenerateKeyMustHaveCorrectKeySize(t *testing.T) {
	size := 8
	data := generateKey(true, true, true, true, size)

	has_symbol := false
	has_lower := false
	has_upper := false
	has_number := false

	for _, c := range data {
		if unicode.IsPunct(c) {
			has_symbol = true
		}
		if unicode.IsDigit(c) {
			has_number = true
		}
		if unicode.IsUpper(c) {
			has_upper = true
		}
		if unicode.IsLower(c) {
			has_lower = true
		}
	}

	assert.True(t, has_symbol)
	assert.True(t, has_lower)
	assert.True(t, has_upper)
	assert.True(t, has_number)
	assert.Equal(t, size, len(data))
}
