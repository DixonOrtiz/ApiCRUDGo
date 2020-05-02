package main

import (
	"fmt"

	"github.com/DixonOrtiz/ApiRap/api"
)

//El paquete main debe importar el paquete api y ejecutar la funcion que corre todo
func main() {
	fmt.Println("Corriendo en el puerto 8000")
	api.Run()
}
