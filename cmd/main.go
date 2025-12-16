package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"prodmod/internal/model"
	"prodmod/internal/repository"
	idpkg "prodmod/pkg/id"
)

func main() {
	repo, err := repository.NewRepository("data/facts.txt", "data/rules.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("facts: ", len(repo.Facts))
	fmt.Println("rules: ", len(repo.Rules))

	axioms, err := idpkg.ReadIds("data/axioms.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("axioms: ", len(axioms))

	targets, err := idpkg.ReadIds("data/target.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("target: ", len(targets))

	fmt.Println("\nИсходные данные (axioms):")
	for _, axiom := range axioms {
		if fact, ok := repo.Facts[axiom]; ok {
			fmt.Printf("- %s (ID: f%02d)\n", fact.Description, fact.Id)
		}
	}

	fmt.Println("\nЦели для проверки:")
	for _, target := range targets {
		if fact, ok := repo.Facts[target]; ok {
			fmt.Printf("- %s (ID: f%02d)\n", fact.Description, fact.Id)
		}
	}

	fmt.Print("\nВыберите режим работы:\n")
	fmt.Println("1. Прямой вывод (от данных к целям)")
	fmt.Println("2. Обратный вывод (доказательство целей)")
	fmt.Print("Введите 1 или 2: ")

	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	m := model.NewProductionModel(axioms, targets, repo)

	switch choice {
	case "1":
		fmt.Println("\n=== РЕЖИМ ПРЯМОГО ВЫВОДА ===")
		m.Run()
		m.Print()
		m.GetAdvice()
		
	case "2":
		fmt.Println("\n=== РЕЖИМ ОБРАТНОГО ВЫВОДА ===")
		results := m.RunBackward()
		m.Print()
		m.GetAdviceBackward(results)
		
	default:
		fmt.Println("Такой вариант не реализован. Запускаю оба режима.")
		
		fmt.Println("\n=== Forward ===")
		m1 := model.NewProductionModel(axioms, targets, repo)
		m1.Run()
		m1.Print()
		m1.GetAdvice()
		
		fmt.Println("\n=== Backward ===")
		m2 := model.NewProductionModel(axioms, targets, repo)
		results := m2.RunBackward()
		m2.Print()
		m2.GetAdviceBackward(results)
	}
}