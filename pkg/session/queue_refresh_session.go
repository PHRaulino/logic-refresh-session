package session

import (
	"fmt"
	"os"
	"time"
)

func AddSessionTime(filename string, delay time.Time) {
	const layout = "2006-01-02 15:04:05"

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo: %s\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(delay.Format(layout) + "\n")
	if err != nil {
		fmt.Printf("Erro ao escrever no arquivo: %s\n", err)
		return
	}
}

func ClearSession(filename string) {
	err := os.Truncate(filename, 0)
	if err != nil {
		fmt.Println("Erro ao truncar o arquivo:", err)
		return
	}
}
