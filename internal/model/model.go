package model

import (
	"prodmod/internal/fact"
)

type ProductionModel struct {
	Memory map[fact.ID]bool
	Target fact.ID
}

func NewProductionModel(axioms []fact.ID, target fact.ID) *ProductionModel {
	memory := make(map[fact.ID]bool)
	for _, axiom := range axioms {
		memory[axiom] = true
	}

	return &ProductionModel{
		Memory: memory,
		Target: target,
	}
}
