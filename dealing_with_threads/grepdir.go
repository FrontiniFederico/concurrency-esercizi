package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	args := os.Args
	for _, arg := range args[2:] {
		grepDir(arg, args[1]) 
	}
}

func grepDir(dir_path string, match string) {
	fileInfo, err := os.ReadDir(dir_path)
	if err != nil {
		fmt.Println("Errore nell'apertura della director:", err)
		return
	}
	// Itera su ogni file nella directory e stampa il nome del file
	fmt.Println("Elenco dei file nella directory:")
	for _, file := range fileInfo {
		// Stampa il nome del file
		fmt.Println(file.Name())
		// file è un oggetto direntry, non una stringa contenente il path
		filePath := filepath.Join(dir_path, file.Name())
		go grep(filePath, match)
		time.Sleep(1)
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
}
