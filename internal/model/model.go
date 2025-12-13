package model

import (
	"fmt"
	"prodmod/internal/fact"
	"prodmod/internal/repository"
)

type ProductionModel struct {
	Memory     map[fact.ID]bool
	Target     fact.ID
	Repository *repository.Repository
}

func NewProductionModel(axioms []fact.ID, target fact.ID, repository *repository.Repository) *ProductionModel {
	memory := make(map[fact.ID]bool)
	for _, axiom := range axioms {
		memory[axiom] = true
	}

	return &ProductionModel{
		Memory:     memory,
		Target:     target,
		Repository: repository,
	}
}

func (p *ProductionModel) ProcessRule(from []fact.ID, to fact.ID) bool {
	if _, exists := p.Memory[to]; exists {
		return false
	}

	for _, factID := range from {
		if _, exists := p.Memory[factID]; !exists {
			return false
		}
	}

	p.Memory[to] = true
	return true
}

func (p *ProductionModel) Run() (bool, error) {
	for {
		factAdded := false

		for _, rule := range p.Repository.Rules {
			if p.ProcessRule(rule.From, rule.Result) {
				factAdded = true
			}
		}
		if _, ok := p.Memory[p.Target]; ok {
			return true, nil
		}
		if !factAdded {
			break
		}
	}
	return false, nil
}

func (p *ProductionModel) Print() {
	fmt.Println("--- Final Memory State ---")
	for id := range p.Memory {
		if f, ok := p.Repository.Facts[id]; ok {
			fmt.Printf("[%d] %s (Weight: %.1f)\n", f.Id, f.Description, f.Weight)
		} else {
			fmt.Printf("[%d] Unknown Fact\n", id)
		}
	}
	fmt.Println("--------------------------")
}
