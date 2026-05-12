package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const fileName = "logfile.log"

	lineNumber := 0
	errorCount := 0

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, " Ошибка открытия %s: %v\n", fileName, err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		if strings.Contains(strings.ToLower(line), "error") {
			errorCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка чтения файла: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Обработано строк: %d\n", lineNumber)
	fmt.Printf(" Найдено строк со словом 'error': %d\n", errorCount)
}
