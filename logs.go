package main

import (
	"errors"
	"log"
	"os"
)

func logs() {
	err := errors.New("Error de prueba")
	log.Println("Error:", err)
	f, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Printf("Error linea %v", 1)
	/*
		log.Panicln("Error:", err)
		log.Panic("Error:", err)
		log.Fatalf("Error: %v", err)
		log.Fatalln(err)
	*/
}
