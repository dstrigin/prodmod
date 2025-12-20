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

	fmt.Println("facts.txt: ", len(repo.Facts))
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

	m := model.NewProductionModel(axioms, targets, repo)

	//m.Run()
	m.ReverseRun()
	m.GetAdvice()
}
