package main

import (
	"bufio"
	"os"
	"fmt"
	"net/rpc"
	
)
type Student struct{
	Nombre string
	Materia string
	Calificacion float64
}
var scanner = bufio.NewScanner(os.Stdin)
func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op int64
	var st Student
	
	for {
		fmt.Println("1 Agregar ")
		fmt.Println("2 promedio por estudiante")
		fmt.Println("3 promedio general")
		fmt.Println("4 promedio por materia")
		fmt.Println("0 Salir")
		fmt.Print("Opcion: ")
		fmt.Scanln(&op)

		switch op {
		case 1:
			
			var calificacion float64
			
			fmt.Print("\nNombre: ")
			scanner.Scan()
			st.Nombre = scanner.Text()
			
			fmt.Print("\nMateria: ")
			scanner.Scan()
			st.Materia = scanner.Text()

			fmt.Print("\nCalificacion: ")
			fmt.Scanln(&calificacion)
			st.Calificacion = calificacion

			

			var result string
			err = c.Call("Servery.Insertar", st, &result)
			if err != nil {
				fmt.Println(err)
			} 

		case 2:

			var nom string
			fmt.Print("\nNombre: ")
			fmt.Scanln(&nom)

			var result string
			err = c.Call("Servery.PromedioSt", nom, &result)
			if err != nil {
				fmt.Println(err)
			} 

		case 3:

			aux:= float64(0)
			var result string
			err = c.Call("Servery.PromedioGSt",  aux, &result)
			
			if err != nil {
				fmt.Println(err)
			} 

		case 4:

			var mat string
			var result string
			fmt.Print("\nMateria: ")
			fmt.Scanln(&mat)
			
			err = c.Call("Servery.PromedioMa",  mat, &result)
			if err != nil {
				fmt.Println(err)
			} 

		case 0:
			return
		}
	}
}

func main() {
	client()
}