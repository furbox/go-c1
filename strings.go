package main

import (
	"fmt"
	"strings"
)

// modulo strings
func stringsx() {
	cadena := "Hola mundo ñandú"
	fmt.Println(cadena)
	fmt.Println(strings.ToUpper(cadena))
	fmt.Println(strings.ToLower(cadena))
	letras := strings.Split(cadena, "")
	fmt.Println(letras)
	posicion := strings.Index(cadena, "mundo")
	if posicion == -1 {
		fmt.Println("No se encontro la palabra mundo")
	} else {
		fmt.Println("La palabra mundo esta en la posicion", posicion)
	}
	repetida := strings.Repeat(cadena, 2)
	fmt.Println(repetida)
	cadenanueva := strings.Replace(cadena, "mundo", "gente", -1)
	fmt.Println(cadenanueva)
	fmt.Println(string(cadenanueva[0:5]))
}
