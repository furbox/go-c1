package main

import (
	"math/rand"
	"strings"
	"time"
)

// math/rand is a package that provides pseudorandom number generation.
func mathrand() {
	aleatorio := rand.Intn(100)
	println("Número aleatório:", aleatorio)
	min := 1000
	max := 10000
	rand.Seed(time.Now().UnixNano())
	aleatorio2 := rand.Intn(max-min) + min
	println("Número aleatório entre 1000 e 10000:", aleatorio2)
	password := generatePassword(20, 1, 1, 1)
	println("Contraseña:", password)
}

var (
	lowerCharSet   = "abcdedfghijklmñnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
