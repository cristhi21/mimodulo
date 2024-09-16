package main

import (
	"mimodulo/videogame"
	"fmt"
)

func main() {
	// Llamada a la función definida en operaciones.go
	resultadoSuma := Sumar(3, 4)
	fmt.Println("Resultado de la suma:", resultadoSuma)

	// Llamada a la función definida en utilidades.go
	saludo := Saludar("ISA")
	fmt.Println(saludo)

	game := videogame.Videogame{
        Id:    1,
        Title: "The Legend of Zelda",
    }

    // Imprimir los detalles del videojuego
    fmt.Println("ID del videojuego:", game.Id)
    fmt.Println("Título del videojuego:", game.Title)
}