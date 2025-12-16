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

func (p *ProductionModel) BackwardChain(target fact.ID, depth int) (bool, []string) {
	var proof []string
	indent := ""
	for i := 0; i < depth; i++ {
		indent += "  "
	}

	if _, known := p.Memory[target]; known {
		desc := p.Repository.Facts[target].Description
		proof = append(proof, fmt.Sprintf("%s✓ %s (ID: f%02d) уже известно", indent, desc, target))
		return true, proof
	}

	foundRules := make([]*rule.Rule, 0)
	for _, rule := range p.Repository.Rules {
		if rule.Result == target {
			foundRules = append(foundRules, rule)
		}
	}

	if len(foundRules) == 0 {
		desc := p.Repository.Facts[target].Description
		proof = append(proof, fmt.Sprintf("%s✗ %s (ID: f%02d) - нет правил для вывода", indent, desc, target))
		return false, proof
	}

	// Пытаемся применить каждое правило
	for _, rule := range foundRules {
		desc := p.Repository.Facts[target].Description
		proof = append(proof, fmt.Sprintf("%sПытаемся доказать %s (ID: f%02d) через правило:", 
			indent, desc, target))
		proof = append(proof, fmt.Sprintf("%s  %s", indent, rule.Description))
		
		allPreconditionsMet := true
		var subProof []string
		
		// Проверяем все предпосылки правила
		for _, precondition := range rule.From {
			preconditionDesc := p.Repository.Facts[precondition].Description
			proof = append(proof, fmt.Sprintf("%s  Проверяем предпосылку: %s (ID: f%02d)", 
				indent, preconditionDesc, precondition))
			
			preconditionMet, preconditionProof := p.BackwardChain(precondition, depth+2)
			subProof = append(subProof, preconditionProof...)
			
			if !preconditionMet {
				allPreconditionsMet = false
				proof = append(proof, fmt.Sprintf("%s    ✗ Предпосылка не доказана", indent))
				break
			} else {
				proof = append(proof, fmt.Sprintf("%s    ✓ Предпосылка доказана", indent))
			}
		}
		
		if allPreconditionsMet {
			// Все предпосылки выполнены, добавляем результат в память
			p.Memory[target] = true
			p.Derivation[target] = rule
			proof = append(proof, subProof...)
			proof = append(proof, fmt.Sprintf("%s✓ %s (ID: f%02d) успешно доказано!", 
				indent, desc, target))
			return true, proof
		}
		
		proof = append(proof, subProof...)
	}
	
	desc := p.Repository.Facts[target].Description
	proof = append(proof, fmt.Sprintf("%s✗ %s (ID: f%02d) - не удалось доказать", 
		indent, desc, target))
	return false, proof
}

func (p *ProductionModel) RunBackward() map[fact.ID]bool {
	results := make(map[fact.ID]bool)
	
	fmt.Println("\n=== ОБРАТНЫЙ ВЫВОД (ДОКАЗАТЕЛЬСТВО ЦЕЛЕЙ) ===")
	
	for targetID := range p.Targets {
		fmt.Printf("\nПытаемся доказать цель: %s (ID: f%02d)\n", 
			p.Repository.Facts[targetID].Description, targetID)
		fmt.Println("----------------------------------------")
		
		// Делаем копию памяти для каждого доказательства
		originalMemory := make(map[fact.ID]bool)
		for k, v := range p.Memory {
			originalMemory[k] = v
		}
		
		originalDerivation := make(map[fact.ID]*rule.Rule)
		for k, v := range p.Derivation {
			originalDerivation[k] = v
		}
		
		proved, proof := p.BackwardChain(targetID, 0)
		results[targetID] = proved
		
		for _, line := range proof {
			fmt.Println(line)
		}
		
		if proved {
			fmt.Printf("\n✓ Цель f%02d доказана!\n", targetID)
		} else {
			fmt.Printf("\n✗ Цель f%02d не доказана\n", targetID)
			// Восстанавливаем исходное состояние памяти
			p.Memory = originalMemory
			p.Derivation = originalDerivation
		}
	}
	
	return results
}

func (p *ProductionModel) GetAdvice() {
	fmt.Println("\n==============================")
	fmt.Println("Recommendations (forward):")

	foundAny := false
	for targetID := range p.Targets {
		if _, known := p.Memory[targetID]; known {
			foundAny = true
			factDesc := p.Repository.Facts[targetID].Description
			fmt.Printf("Recommendations: %s (ID: f%02d)\n", factDesc, targetID)

			if r, ok := p.Derivation[targetID]; ok {
				fmt.Printf("Target: %s\n", r.Description)
			}
		}
	}

	if !foundAny {
		fmt.Println("К сожалению, на основе введенных данных конкретных рекомендаций нет.")
	}
	fmt.Println("==============================")
}

func (p *ProductionModel) GetAdviceBackward(results map[fact.ID]bool) {
	fmt.Println("\n==============================")
	fmt.Println("РЕКОМЕНДАЦИИ (backward):")

	foundAny := false
	for targetID, proved := range results {
		if proved {
			foundAny = true
			factDesc := p.Repository.Facts[targetID].Description
			fmt.Printf("Рекомендация: %s (ID: f%02d)\n", factDesc, targetID)

			if r, ok := p.Derivation[targetID]; ok {
				fmt.Printf("\tПричина: %s\n", r.Description)
			}
		}
	}

	if !foundAny {
		fmt.Println("Не удалось доказать ни одну из целей.")
	}
	fmt.Println("==============================")
}

func (p *ProductionModel) Print() {
	fmt.Println("\n--- Final Memory State ---")
	for id := range p.Memory {
		if f, ok := p.Repository.Facts[id]; ok {
			fmt.Printf("[f%02d] %s (%.2f)\n", f.Id, f.Description, f.Weight)
		} else {
			fmt.Printf("[f%02d] Unknown Fact\n", id)
		}
	}
	fmt.Println("--------------------------")
}