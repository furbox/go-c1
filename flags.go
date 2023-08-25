package main

import "flag"

// modulo os | argumentos de entrada
func flags() {
	//go run main.go -nombre Jorge -edad 22
	nombre := flag.String("nombre", "", "Nombre del usuario")
	edad := flag.Int("edad", 18, "Edad del usuario")
	flag.Parse()
	println("Nombre:", *nombre)
	println("Edad:", *edad)
}
