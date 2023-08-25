package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// funcion para conectar a la base de datos
var Db *sql.DB

func Conectar() {
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic("Error al cargar las variables de entorno")
	}
	db, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@tcp("+os.Getenv("DB_SERVER")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}
	Db = db
}

// cerrar Conexion
func Desconectar() {
	Db.Close()
}

//go get github.com/go-sql-driver/mysql
//go get github.com/joho/godotenv

/*
CREATE TABLE `clientes` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `nombre` varchar(100) NOT NULL,
  `correo` varchar(100) NOT NULL,
  `telefono` varchar(20) NOT NULL,
  `fecha` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='clientes'
*/
