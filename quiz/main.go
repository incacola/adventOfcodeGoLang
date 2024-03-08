package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	var respuestasCorrectas int
	var respuestasIncorrectas int
	var sliceDePreguntas []string
	var sliceDeRespuestas []int
	var archivoDeProblemas = lectorDeProblemas("./problems.csv")

	fmt.Printf("Hola bienvenido a tu Quiz responde lo mas rapido posible!\n")
	for u := 0; u < len(archivoDeProblemas); u++ {
		lineaDePregunta := strings.Split(archivoDeProblemas[u], ",")
		sliceDePreguntas = append(sliceDePreguntas, lineaDePregunta[0])

		num, e := strconv.Atoi(lineaDePregunta[1])
		if e == nil {
			sliceDeRespuestas = append(sliceDeRespuestas, int(num))
		} else {
			panic(e)
		}
	}

	var cantidadDePreguntas int = rand.IntN(99) + 1
	for i := 0; i < cantidadDePreguntas; i++ {
		var preguntaAliatoria int = rand.IntN(11) + 1
		var pregunta = sliceDePreguntas[preguntaAliatoria]

		var respuestaCorrecta = sliceDeRespuestas[preguntaAliatoria]

		var respuesta int
		fmt.Printf("Cuanto es: %v \n", pregunta)
		fmt.Scan(&respuesta)

		if respuesta == respuestaCorrecta {
			respuestasCorrectas++
		} else {
			respuestasIncorrectas++
		}
	}

	fmt.Printf("Esta son la cantidad de respuestas correctas %v \n", respuestasCorrectas)
	fmt.Printf("Esta son la cantidad de respuestas incorrectas %v \n", respuestasIncorrectas)
}

func lectorDeProblemas(path string) []string {
	dat, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(dat)
	scanner.Split(bufio.ScanLines)
	var lineasDeTexto []string

	for scanner.Scan() {
		lineasDeTexto = append(lineasDeTexto, scanner.Text())
	}

	return lineasDeTexto
}
