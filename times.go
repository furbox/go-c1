package main

import (
	"fmt"
	"time"
)

// times
func times() {
	//obtener la fecha y hora actual del servidor
	fecha := time.Now()
	//imprimir la fecha y hora
	fmt.Println(FormatearFechaHora(fecha))
	ahora := time.Now()
	fmt.Println(FormatearFecha(ahora))
	//agregar 20 dias mas
	fecha1 := ahora.Add(time.Hour * 24 * 20)
	fmt.Println(FormatearFecha(fecha1))
	//restar 20 dias a ahora
	fecha2 := ahora.Add(time.Hour * 24 * -20)
	fmt.Println(FormatearFecha(fecha2))
}

func FormatearFecha(fecha time.Time) string {
	v := fmt.Sprintf("%v/%v/%v", fecha.Day(), int(fecha.Month()), fecha.Year())
	return v
}

func FormatearFechaHora(fecha time.Time) string {
	v := fmt.Sprintf("%v/%v/%v %v:%v:%v", fecha.Day(), int(fecha.Month()), fecha.Year(), fecha.Hour(), fecha.Minute(), fecha.Second())
	return v
}
