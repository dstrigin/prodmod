package repository

import (
	"prodmod/internal/fact"
	"prodmod/internal/rule"
)

type ReverseRepository struct {
	Producers map[fact.ID][]*rule.Rule
}

func NewReverseRepository(repo *Repository) *ReverseRepository {
	producers := make(map[fact.ID][]*rule.Rule)
	for _, r := range repo.Rules {
		producers[r.Result] = append(producers[r.Result], r)
	}
	return &ReverseRepository{Producers: producers}
}
