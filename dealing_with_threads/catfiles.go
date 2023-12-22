package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args
	for _, arg := range args[1:] {
		go cat(arg) //solo questa fa terminare tutto
		// runtime.Gosched() //aggiungere questa OGNI TANTO ne fa stampare uno
		time.Sleep(1) //va meglio, ma OGNI TANTO ne stampa in meno
	}
}

func cat(file_path string) {
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
		fmt.Println(linea)
	}

	// Controlla errori durante la scansione del file
	if err := scanner.Err(); err != nil {
		fmt.Println("Errore durante la lettura del file:", err)
	}
}