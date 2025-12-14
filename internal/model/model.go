package model

import (
	"fmt"
	"prodmod/internal/fact"
	"prodmod/internal/repository"
	"prodmod/internal/rule"
)

type ProductionModel struct {
	Memory     map[fact.ID]bool
	Targets    map[fact.ID]bool
	Repository *repository.Repository
	Derivation map[fact.ID]*rule.Rule
}

func NewProductionModel(axioms []fact.ID, targets []fact.ID, repository *repository.Repository) *ProductionModel {
	memory := make(map[fact.ID]bool)
	for _, axiom := range axioms {
		memory[axiom] = true
	}

	targetMap := make(map[fact.ID]bool)
	for _, target := range targets {
		targetMap[target] = true
	}

	return &ProductionModel{
		Memory:     memory,
		Targets:    targetMap,
		Repository: repository,
		Derivation: make(map[fact.ID]*rule.Rule),
	}
}

func (p *ProductionModel) ProcessRule(r *rule.Rule) bool {
	if _, exists := p.Memory[r.Result]; exists {
		return false
	}

	for _, factID := range r.From {
		if _, exists := p.Memory[factID]; !exists {
			return false
		}
	}

	p.Memory[r.Result] = true
	p.Derivation[r.Result] = r
	return true
}

func (p *ProductionModel) Run() {
	for {
		factAdded := false

		for _, rule := range p.Repository.Rules {
			if p.ProcessRule(rule) {
				factAdded = true
			}
		}
		if !factAdded {
			break
		}
	}
}

func (p *ProductionModel) GetAdvice() {
	fmt.Println("==============================")

	foundAny := false
	for targetID := range p.Targets {
		if _, known := p.Memory[targetID]; known {
			foundAny = true
			factDesc := p.Repository.Facts[targetID].Description
			fmt.Println("Рекомендация:", factDesc)

			if r, ok := p.Derivation[targetID]; ok {
				fmt.Println("\tПричина:", r.Description)
			}
		}
	}

	if !foundAny {
		fmt.Println("К сожалению, на основе введенных данных конкретных рекомендаций нет.")
	}
	fmt.Println("==============================")
}

func (p *ProductionModel) Print() {
	fmt.Println("--- Final Memory State ---")
	for id := range p.Memory {
		if f, ok := p.Repository.Facts[id]; ok {
			fmt.Printf("[%v] %s (%v)\n", f.Id, f.Description, f.Weight)
		} else {
			fmt.Printf("[%v] Unknown Fact\n", id)
		}
	}
	fmt.Println("--------------------------")
}
