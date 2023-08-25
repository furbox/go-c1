package handlers

import (
	"bufio"
	"fmt"
	"go_c1/db"
	"go_c1/modelos"
	"os"
	"strconv"
)

// listar clientes
func Listar() {
	db.Conectar()
	sql := "SELECT * FROM clientes ORDER BY id DESC;"
	datos, err := db.Db.Query(sql)
	if err != nil {
		panic(err)
	}
	db.Desconectar()
	fmt.Println("\nListado de Clientes")
	fmt.Println("---------------------------------")
	for datos.Next() {

		var id int
		var nombre string
		var correo string
		var telefono string
		err := datos.Scan(&id, &nombre, &correo, &telefono)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Nombre:", nombre)
		fmt.Println("Correo:", correo)
		fmt.Println("Telefono:", telefono)
		fmt.Println("---------------------------------")
	}

}

// Listar por id
func ListarPorId(id int) {
	db.Conectar()
	sql := "SELECT * FROM clientes WHERE id = ?;"
	datos, err := db.Db.Query(sql, id)
	if err != nil {
		panic(err)
	}
	db.Desconectar()
	fmt.Println("\nCliente")
	fmt.Println("---------------------------------")
	for datos.Next() {

		var id int
		var nombre string
		var correo string
		var telefono string
		err := datos.Scan(&id, &nombre, &correo, &telefono)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Nombre:", nombre)
		fmt.Println("Correo:", correo)
		fmt.Println("Telefono:", telefono)
		fmt.Println("---------------------------------")
	}

}

// Insertar cliente
func Insertar(cliente modelos.Cliente) {
	db.Conectar()
	sql := "INSERT INTO clientes(nombre, correo, telefono) VALUES(?,?,?);"
	result, err := db.Db.Exec(sql, cliente.Nombre, cliente.Correo, cliente.Telefono)
	db.Desconectar()
	if err != nil {
		panic(err)
	}
	//obtener el ultimo id insertado
	id, _ := result.LastInsertId()
	fmt.Println("Id insertado:", id)
	fmt.Println("Cliente insertado correctamente")
}

// editar cliente
func Editar(id int, cliente modelos.Cliente) {
	db.Conectar()
	sql := "UPDATE clientes SET nombre = ?, correo = ?, telefono = ? WHERE id = ?;"
	result, err := db.Db.Exec(sql, cliente.Nombre, cliente.Correo, cliente.Telefono, id)
	db.Desconectar()
	if err != nil {
		panic(err)
	}
	//obtener el numero de filas afectadas
	numeroFilas, _ := result.RowsAffected()
	fmt.Println("Numero de filas afectadas:", numeroFilas)
	fmt.Println("Cliente actualizado correctamente")
}

// eliminar cliente
func Eliminar(id int) {
	db.Conectar()
	sql := "DELETE FROM clientes WHERE id = ?;"
	result, err := db.Db.Exec(sql, id)
	db.Desconectar()
	if err != nil {
		panic(err)
	}
	//obtener el numero de filas afectadas
	numeroFilas, _ := result.RowsAffected()
	fmt.Println("Numero de filas afectadas:", numeroFilas)
	fmt.Println("Cliente eliminado correctamente")
}

// funcion ejecutar
var ID int
var nombre, correo, telefono string

func Ejecutar() {
	scanner := bufio.NewScanner(os.Stdin)
	println("Seleccione una Opción: \n\n")
	println("1. Listar Clientes")
	println("2. Listar Cliente por ID")
	println("3. Insertar Cliente")
	println("4. Editar Cliente")
	println("5. Eliminar Cliente")
	if scanner.Scan() {
		for {
			if scanner.Text() == "1" {
				Listar()
				Ejecutar()
				return
			}
			if scanner.Text() == "2" {
				println("Ingrese el ID del Cliente")
				if scanner.Scan() {
					ID, _ = strconv.Atoi(scanner.Text())
					ListarPorId(ID)
					Ejecutar()
					return
				}
			}
			if scanner.Text() == "3" {
				println("Ingrese el Nombre del Cliente")
				if scanner.Scan() {
					nombre = scanner.Text()
				}
				println("Ingrese el Correo del Cliente")
				if scanner.Scan() {
					correo = scanner.Text()
				}
				println("Ingrese el Telefono del Cliente")
				if scanner.Scan() {
					telefono = scanner.Text()
				}
				cliente := modelos.Cliente{
					Nombre:   nombre,
					Correo:   correo,
					Telefono: telefono,
				}
				Insertar(cliente)
				Ejecutar()
				return
			}
			if scanner.Text() == "4" {
				println("Ingrese el ID del Cliente")
				if scanner.Scan() {
					ID, _ = strconv.Atoi(scanner.Text())
				}
				println("Ingrese el Nombre del Cliente")
				if scanner.Scan() {
					nombre = scanner.Text()
				}
				println("Ingrese el Correo del Cliente")
				if scanner.Scan() {
					correo = scanner.Text()
				}
				println("Ingrese el Telefono del Cliente")
				if scanner.Scan() {
					telefono = scanner.Text()
				}
				cliente := modelos.Cliente{
					Nombre:   nombre,
					Correo:   correo,
					Telefono: telefono,
				}
				Editar(ID, cliente)
				Ejecutar()
				return
			}
			if scanner.Text() == "5" {
				println("Ingrese el ID del Cliente")
				if scanner.Scan() {
					ID, _ = strconv.Atoi(scanner.Text())
				}
				Eliminar(ID)
				Ejecutar()
				return
			}
		}
	}
}
