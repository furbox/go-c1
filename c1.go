package main

import "fmt"

// interfaces
func mainx() {
	e := Estructura{}
	fmt.Println(e.miFuncion())
	fmt.Println(e.Calculo(5, 7))
}

type EjemploInterface interface {
	miFuncion() string
	Calculo(n1 int, n2 int) int
}

type Estructura struct{}

func (*Estructura) miFuncion() string {
	return "Hola mundo"
}

func (*Estructura) Calculo(n1 int, n2 int) int {
	return n1 + n2
}

/*
// estructuras
func main() {
	estructura := Persona{
		Id:     1,
		Nombre: "Chris",
		Correo: "furbox@gmail.com",
		Edad:   37,
	}
	fmt.Println(estructura)
	fmt.Println(reflect.TypeOf(estructura))
	fmt.Println(estructura.Nombre)
	p := new(Persona)
	fmt.Println(reflect.TypeOf(p))
	p.Id = 2
	p.Nombre = "Flores"
	p.Correo = "furbox@gmail.com"
	p.Edad = 37
	fmt.Println(p)

	categoria := Categotia{
		Id:     1,
		Nombre: "Ropa",
		Slug:   "ropa",
	}
	producto := Producto{
		Id:          1,
		Nombre:      "Camisa",
		Descripcion: "Camisa de algodon",
		Slug:        "camisa",
		CategoriaId: categoria,
	}
	fmt.Println(producto)
}

type Persona struct {
	Id     int
	Nombre string
	Correo string
	Edad   int
}

type Categotia struct {
	Id     int
	Nombre string
	Slug   string
}

type Producto struct {
	Id          int
	Nombre      string
	Descripcion string
	Slug        string
	CategoriaId Categotia
}

/*
// defer y panic
func main() {
	miFuncion()
}

func miFuncion() {
	defer fmt.Println("Esto se ejecuta al final")
	fmt.Println("Hola mundo")
	a := 1
	if a == 1 {
		panic("Se encontro el valor de 1")
	}
}

/*
// recursividad
func main() {
	miFuncion(0)
}

func miFuncion(valor int) {
	dato := valor + 1
	fmt.Println(dato)
	if dato < 10 {
		miFuncion(dato)
	}
}

/*
// goroutines y channels
func main() {
	//ejemplo artesanal
	fmt.Println(miFuncion("Chris"))
	time.Sleep(time.Second * 5)
	fmt.Println(miFuncion("Flores"))

	//channels
	channel := make(chan string)
	go func() {
		channel <- miFuncion("Chris")
	}()
	fmt.Println(<-channel)
	fmt.Println("continuamos con el flujo de la aplicacion")

}

func miFuncion(nombre string) string {
	return "Hola " + nombre
}

/*
// funciones
func main() {
	miFuncion()
	miFuncionConParametros(10, 20)
	fmt.Println(miFuncionConRetorno("Chris"))
	nombre, apellido, edad := miFuncionConRetornoMultiple("Chris")
	fmt.Println(nombre, apellido, edad)
	fmt.Println(suma(10, 20))
	// funciones closures
	Tabla := tabla(2)
	for i := 1; i <= 10; i++ {
		fmt.Printf("2 x %v = %v \n", i, Tabla())
	}
}

func miFuncion() {
	fmt.Println("Mi primera funcion")
}

func miFuncionConParametros(n1 int, n2 int) {
	resultado := n1 + n2
	fmt.Println("El resultado de la suma es: ", resultado)
}

func miFuncionConRetorno(nombre string) string {
	return "Tu nombres es: " + nombre
}

func miFuncionConRetornoMultiple(nombre string) (string, string, int) {
	return "Chris", "Flores", 37
}

// funcion anonima
var suma = func(n1 int, n2 int) int {
	return n1 + n2
}

// funciones closures
func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

/*
// map
func main() {
	//{"id":1,"nombre":"Mexico"}
	paises := make(map[string]int)
	//pais y habitantes
	paises["Mexico"] = 120000000
	paises["Argentina"] = 45000000
	paises["Colombia"] = 50000000
	paises["Peru"] = 33000000
	fmt.Println(paises)
	fmt.Println(paises["Mexico"])
	fmt.Println(len(paises))
	fmt.Println(reflect.TypeOf(paises))
	//eliminar un elemento del map
	delete(paises, "Peru")
	fmt.Println(paises)
	fmt.Println(paises["Mexico"])
	fmt.Println(len(paises))
	//recorrer un map
	for pais, habitantes := range paises {
		fmt.Println(pais, habitantes)
	}
	//otro formato de map
	paises2 := map[string]int{
		"Mexico":    120000000,
		"Argentina": 45000000,
		"Colombia":  50000000,
		"Peru":      33000000,
	}
	fmt.Println(paises2)
	fmt.Println(paises2["Mexico"])
	fmt.Println(len(paises2))
	fmt.Println(reflect.TypeOf(paises2))
	//validar si existe algun elemento en el map
	valor, existe := paises2["Peru"]
	if existe {
		fmt.Println(valor, "Existe")
	} else {
		fmt.Println(valor, "No existe")
	}
	//mensaje de respuesta
	respuesta := map[string]string{
		"estado": "ok",
		"msg":    "mensaje de respuesta",
	}
	fmt.Println(respuesta)
}

/*
// arreglos y slices
func main() {
	//arreglo (array)
	var paises [4]string
	paises[0] = "Colombia"
	paises[1] = "Mexico"
	paises[2] = "Argentina"
	paises[3] = "Peru"
	fmt.Println(paises)
	fmt.Println(paises[0])
	fmt.Println(len(paises))
	fmt.Println(reflect.TypeOf(paises))

	//slice
	paises2 := make([]string, 5)
	paises2[0] = "Colombia"
	paises2[1] = "Mexico"
	paises2[2] = "Argentina"
	paises2[3] = "Peru"
	paises2[4] = "España"
	fmt.Println(paises2)
	fmt.Println(paises2[0])
	fmt.Println(len(paises2))
	fmt.Println(reflect.TypeOf(paises2))
	//agregar un elemento al slice
	paises2 = append(paises2, "Alemania")
	fmt.Println(paises2)
	fmt.Println(paises2[0])
	fmt.Println(len(paises2))
	fmt.Println(reflect.TypeOf(paises2))
	//eliminar un elemento del slice
	paises2 = append(paises2[:2], paises2[3:]...)
	fmt.Println(paises2)
	fmt.Println(paises2[0])
	fmt.Println(len(paises2))
	fmt.Println(reflect.TypeOf(paises2))
}


// ciclos e iteraciones
func main() {
	i := 1
	for i < 11 {
		fmt.Println(i)
		i++
	}

	for i2 := 1; i2 < 11; i2++ {
		fmt.Println(i2)
	}

	//imprimir numeros pares
	for i3 := 1; i3 < 11; i3++ {
		if i3%2 == 0 {
			fmt.Println(i3)
		}
	}

	//imprimir numeros impares
	for i4 := 1; i4 < 11; i4++ {
		if i4%2 != 0 {
			fmt.Println(i4)
		}
	}

	//imprimir del 10 al 0
	for i5 := 10; i5 >= 0; i5-- {
		fmt.Println(i5)
	}
}

// condicionales
func main() {
	//operadores de comparacion
	//x == y | x es igual a y
	//x != y | x es diferente a y
	//x < y | x es menor a y
	//x > y | x es mayor a y
	//x <= y | x es menor o igual a y
	//x >= y | x es mayor o igual a y

	edad := 11
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}

	color := "rojo"

	if color == "rojo" {
		fmt.Println("El color es rojo")
	} else if color == "azul" {
		fmt.Println("El color es azul")
	} else {
		fmt.Println("El color no es ni rojo ni azul")
	}

	//operadores logicos &&(and) ||(or) !(not)
	if color == "rojo" && edad >= 18 {
		fmt.Println("Eres mayor de edad y tu color favorito es el rojo")
	} else {
		fmt.Println("No eres mayor de edad y tu color favorito no es el rojo")
	}

	//declarar variables en una condicion
	if variable := 10; variable > 5 {
		fmt.Println("La variable es mayor a 5")
	} else {
		fmt.Println("La variable es menor a 5")
	}

	//switch
	switch color {
	case "rojo":
		fmt.Println("El color es rojo")
	case "azul":
		fmt.Println("El color es azul")
	default:
		fmt.Println("El color no es ni rojo ni azul")
	}
}


// punteros
func main() {
	estado := true
	color := "rojo"
	fmt.Println(color, &color)
	fmt.Println(estado, &estado)
}

// reflect y TypeOf
func main() {

	var string1 string = "Hola"
	fmt.Println(reflect.TypeOf(string1))
}


// tipos de datos
func main() {
	var string1 string = "Hola"
	fmt.Println(string1)

	textoGrande := `Lorem ipsum dolor sit amet, consectetur adipiscing elit.`
	fmt.Println(textoGrande)

	var estado bool = true
	fmt.Println(estado)

	var float32 float32 = 3.14
	fmt.Println(float32)

	var float64 float64 = 3.14
	fmt.Println(float64)

	var entero8 int8 = 127
	fmt.Println(entero8)

	var entero16 int16 = 32767
	fmt.Println(entero16)

	var entero32 int32 = 2147483647
	fmt.Println(entero32)

	var entero64 int64 = 9223372036854775807
	fmt.Println(entero64)

	var entero int = 9223372036854775807
	fmt.Println(entero)

	var enteroSinSigno uint = 18446744073709551615
	fmt.Println(enteroSinSigno)

	var enteroSinSigno8 uint8 = 255
	fmt.Println(enteroSinSigno8)

	var enteroSinSigno16 uint16 = 65535
	fmt.Println(enteroSinSigno16)

	var enteroSinSigno32 uint32 = 4294967295
	fmt.Println(enteroSinSigno32)

	var enteroSinSigno64 uint64 = 18446744073709551615
	fmt.Println(enteroSinSigno64)
}

// variables y constantes
// constantes
const pi float64 = 3.14
const Publica = "Variable publica ñandú"

func main() {
	//declaracion por inferencia
	var nombre string = "Chris"
	fmt.Println(nombre)
	//declaracion rapida o corta
	apellido := "Perez"
	fmt.Println(apellido)

	//imprimir argumentos
	fmt.Println("El valor de mi constante es: ", pi)
	fmt.Printf("El valor de mi variable publica es: %v \n", Publica)
}
*/
