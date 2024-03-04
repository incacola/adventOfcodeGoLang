package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	// Hacer una variable con todo el contenido del archivo
	// calibrations.txt
	calibrationFile := readCalibrationFile("./calibrations.txt")

	// Array con las cuatro lineas del ejemplo
	// Mandar el slice cada linea con un Loop a la function de calibration
	// Poner el resultado en un slice nuevo con los valores
	//
	// calibrationLines := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	calibrationLines := calibrationFile
	calibratedLines := make([]int, 0)

	//Ejecutar funcciones
	for i := 0; i < len(calibrationLines); i++ {
		var calibratedNumber int = calibration(calibrationLines[i])
		calibratedLines = append(calibratedLines, calibratedNumber)
	}
	// calculate value of calibrations
	var resultOrCalibration int = calculateCalibrations(calibratedLines)

	// Prinln del resultado en la consola
	fmt.Printf("Resultado despues de calibracio %v \n", resultOrCalibration)
}

func readCalibrationFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	return txtlines
}

func calculateCalibrations(number []int) int {
	var sum int
	for i := 0; i < len(number); i++ {
		sum = sum + number[i]
	}
	return sum
}

// Function que nos de el resultado de una calibration
// input (1abc) return 12
func calibration(calibrationLine string) int {
	// pasa algo para convertir el String a Int
	var line = calibrationLine

	// Convert the string to a slice of runes
	letters := []rune(line)
	// create the empty numbers slice
	numbers := make([]string, 0)

	// Search numbers and add them to the numbers slice
	for i := 0; i < len(letters); i++ {
		if unicode.IsDigit(letters[i]) {
			var n = string(letters[i])
			numbers = append(numbers, n)
		}
	}

	var returnValue string
	if len(numbers) > 1 {
		returnValue = numbers[0] + numbers[len(numbers)-1]
	} else {
		returnValue = numbers[0] + numbers[0]
	}

	num, e := strconv.Atoi(returnValue)
	if e == nil {
		return num
	}

	return 0
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
