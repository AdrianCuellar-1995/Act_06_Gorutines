package main

import (
	"fmt"
	"time"
)

func Proceso(id uint64) {
	i := uint64(0)
	for {
		if activar {
			fmt.Println(id, " : ", i)
		}
		i = i + 1
		time.Sleep(time.Millisecond * 500)
		if eliminarutina == id {
			eliminarutina = 0
			break
		}
	}
}

var activar bool = false
var eliminarutina uint64 = 0
var contador uint64 = 0

func main() {
	var op int
	for {
		fmt.Println("1) Agregar proceso")
		fmt.Println("2) Mostrar proceso")
		fmt.Println("3) Eliminar procesos")
		fmt.Println("0) Salir")
		fmt.Scanln(&op)

		switch op {
		case 1:
			contador++
			go Proceso(contador)
		case 2:
			if contador > 0 {
				activar = true
				fmt.Scanln()
				activar = false
			} else {
				fmt.Println("No hay procesos registrados")
			}
		case 3:
			if contador > 0 {
				var error bool = true
				for error {
					fmt.Scanln(&eliminarutina)
					if eliminarutina > contador || eliminarutina == 0 {
						fmt.Println("No existe el proceso")
						error = true
					} else {
						error = false
						contador--
					}
				}
			} else {
				fmt.Println("Tienes que registras algun proceso")
			}

		case 0:
			return
		}
	}
}
