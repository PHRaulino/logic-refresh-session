package main

import (
	"fmt"
	"log"
	"time"

	"phraulino/proxy/config"
	"phraulino/proxy/pkg/session"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	startTime := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	stepsRefresh, err := session.GetRefreshSteps(cfg.RefreshTime, startTime, startTime.Add(15*time.Minute))
	if err != nil {
		log.Fatalf("Erro ao recuperar steps: %v", err)
	}

	startTime = time.Date(2024, 1, 1, 0, 4, 0, 0, time.UTC)

	stepsRefreshTwo, err := session.GetRefreshSteps(cfg.RefreshTime, startTime, startTime.Add(15*time.Minute))
	if err != nil {
		log.Fatalf("Erro ao recuperar steps: %v", err)
	}

	existingSteps := make(map[time.Time]bool)
	for _, t := range stepsRefresh {
		existingSteps[t] = true
	}

	for _, t := range stepsRefreshTwo {
		if !existingSteps[t] {
			stepsRefresh = append(stepsRefresh, t)
		}
	}
	startTime = time.Date(2024, 1, 1, 0, 16, 0, 0, time.UTC)

	stepsRefreshThree, err := session.GetRefreshSteps(cfg.RefreshTime, startTime, startTime.Add(15*time.Minute))
	if err != nil {
		log.Fatalf("Erro ao recuperar steps: %v", err)
	}

	existingSteps = make(map[time.Time]bool)
	for _, t := range stepsRefresh {
		existingSteps[t] = true
	}

	for _, t := range stepsRefreshThree {
		if !existingSteps[t] {
			stepsRefresh = append(stepsRefresh, t)
		}
	}

	for _, step := range stepsRefresh {
		fmt.Println(step)
	}
}
