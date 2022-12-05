package pwdgen

import (
	"math/rand"
	"time"
)

var AlphasLower = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
var AlphasUpper = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
var Digits = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
var Specials = []byte{'~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '{', '}', '[', ']'}

const shuffleCount = 50

type config struct {
	UpperAlphas int
	LowerAlphas int
	Digits      int
	Specials    int
}

var Config = config{UpperAlphas: 4, LowerAlphas: 4, Digits: 4, Specials: 4}

func getRandChars(chars []byte, length int) []byte {
	var rand1 = rand.New(rand.NewSource(time.Now().UnixNano()))
	var rchars []byte
	var max = len(chars)

	for i := 0; i < length; i++ {
		rchars = append(rchars, chars[rand1.Intn(max)])
	}
	return rchars
}

func shuffleString(s string) string {
	var rand1 = rand.New(rand.NewSource(time.Now().UnixNano()))
	chars := []byte(s)
	var max = len(chars)

	for i := 1; i < shuffleCount; i++ {
		x := rand1.Intn(max)
		y := rand1.Intn(max)

		chars[x], chars[y] = chars[y], chars[x]
	}
	return string(chars[:])
}

func Password() string {
	lowerAlphaChars := getRandChars(AlphasLower, Config.LowerAlphas)
	upperAlphaChars := getRandChars(AlphasUpper, Config.UpperAlphas)
	digitChars := getRandChars(Digits, Config.Digits)
	specialChars := getRandChars(Specials, Config.Specials)

	allChars := string(lowerAlphaChars[:]) + string(upperAlphaChars[:]) + string(digitChars[:]) + string(specialChars[:])

	return shuffleString(allChars)
}
