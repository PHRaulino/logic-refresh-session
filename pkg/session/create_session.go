package session

import (
	"fmt"
	"os"
	"time"
)

func NewSession() {
	const layout = "2006-01-02 15:04:05"

	now := time.Now()
	future := now.Add(15 * time.Minute)

	futureTime := future.Format(layout)

	filename := "time_record.txt"

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo: %s\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(futureTime + "\n")
	if err != nil {
		fmt.Printf("Erro ao escrever no arquivo: %s\n", err)
		return
	}

	fmt.Println("Data e hora registradas com sucesso:", futureTime)
}
