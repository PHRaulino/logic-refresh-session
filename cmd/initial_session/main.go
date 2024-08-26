package main

import (
	"fmt"
	"log"
	"time"

	"phraulino/proxy/config"
	"phraulino/proxy/pkg/session"
)

func dateInSlice(date time.Time, dates []time.Time) bool {
	for _, d := range dates {
		if d.Equal(date) {
			return true
		}
	}
	return false
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	usersInteractions := []time.Time{
		time.Date(2024, 1, 1, 0, 4, 0, 0, time.UTC),
		time.Date(2024, 1, 1, 0, 5, 0, 0, time.UTC),
		time.Date(2024, 1, 1, 0, 10, 0, 0, time.UTC),
		time.Date(2024, 1, 1, 0, 14, 0, 0, time.UTC),
		time.Date(2024, 1, 1, 0, 24, 0, 0, time.UTC),
		time.Date(2024, 1, 1, 0, 25, 0, 0, time.UTC),
		time.Date(2024, 1, 1, 0, 26, 0, 0, time.UTC),
		time.Date(2024, 1, 1, 0, 28, 0, 0, time.UTC),
		time.Date(2024, 1, 1, 0, 35, 0, 0, time.UTC),
		time.Date(2024, 1, 1, 0, 41, 0, 0, time.UTC),
	}

	var lastSession time.Time
	var fisrtToken time.Time
	var queue []time.Time
	expireToken := usersInteractions[0].Add(5 * time.Minute)
	line := "-------------------------------------------------------------"

	for index, interaction := range usersInteractions {
		fmt.Println(line)
		if index == 0 {
			fmt.Println("Token do usuário criado!")
			fisrtToken = interaction
			continue
		}

		token, err := session.GetTokenSession(cfg.RefreshTime, cfg.SessionTime, interaction, fisrtToken)
		if err != nil {
			log.Fatalf("erro ao recuperar o token")
		}

		if interaction.After(expireToken) {
			panic("Sessão do usuário invalida")
		}

		fmt.Println("Usuário interagiu com o sistema", interaction)
		stepsRefresh, err := session.GetRefreshSteps(cfg.RefreshTime, token, interaction.Add(time.Duration(cfg.SessionTime.Minutes())*time.Minute))
		if err != nil {
			log.Fatalf("Erro ao recuperar steps: %v", err)
		}

		existingSteps := make(map[time.Time]bool)
		for _, t := range queue {
			existingSteps[t] = true
		}

		for _, t := range stepsRefresh {
			if !dateInSlice(t, queue) {
				queue = append(queue, t)
			}
		}

		fmt.Println("Status da Fila")
		for _, step := range queue {
			if step.Before(token) {
				fmt.Println(step, "\t Processado")
			} else if token == step {
				fmt.Println(step, "\t Atual")
			} else {
				fmt.Println(step, "\t A processar")
			}
		}
		lastSession = interaction.Add(time.Duration(cfg.SessionTime.Minutes()) * time.Minute)
		expireToken = queue[len(queue)-1].Add(5 * time.Minute)
		fmt.Println("A sessão do usuario irá expirar em", lastSession)
		fmt.Println("O token do usuário irá expirar de fato", expireToken)
		fmt.Println(line)
	}
}
