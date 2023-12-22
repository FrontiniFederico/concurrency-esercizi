package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"strings"
)

func main() {
	args := os.Args
	for _, arg := range args[2:] {
		go grep(arg, args[1]) //solo questa fa terminare tutto
		// runtime.Gosched() //aggiungere questa OGNI TANTO ne fa stampare uno
		time.Sleep(1) //va meglio, ma OGNI TANTO ne stampa in meno
	}
}

func grep(file_path string, match string) {
	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println("Errore nell'apertura del file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Legge e stampa il contenuto del file riga per riga
	for scanner.Scan() {
		linea := scanner.Text()
		if strings.Contains(linea, match) {
			fmt.Printf("La parola '%s' è presente nella stringa: \"%s\"\n", match, linea)
		}
	}

	// Controlla errori durante la scansione del file
	if err := scanner.Err(); err != nil {
		fmt.Println("Errore durante la lettura del file:", err)
	}
}