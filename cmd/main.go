package main

import (
	"fmt"
	"log"

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

	target, err := idpkg.ReadIds("data/target.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("target: ", len(target))

	m := model.NewProductionModel(axioms, target[0], repo)

	success, err := m.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("production: ", success)

	m.Print()
}
