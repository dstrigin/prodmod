package repository

import (
	"bufio"
	"os"
	"prodmod/internal/fact"
	"prodmod/internal/rule"
)

type Repository struct {
	Facts map[int]*fact.Fact
	Rules map[int]*rule.Rule
}

func NewRepository(factsFileName, rulesFileName string) (*Repository, error) {
	factsFile, err := os.Open(factsFileName)
	if err != nil {
		return nil, err
	}
	defer factsFile.Close()

	facts := make(map[int]*fact.Fact)
	scannerFacts := bufio.NewScanner(factsFile)
	for scannerFacts.Scan() {
		line := scannerFacts.Text()
		fact, err := fact.FromString(line)
		if err != nil {
			return nil, err
		}
		facts[fact.Id] = fact
	}

	rulesFile, err := os.Open(rulesFileName)
	if err != nil {
		return nil, err
	}
	defer rulesFile.Close()

	rules := make(map[int]*rule.Rule)
	scannerRules := bufio.NewScanner(rulesFile)
	for scannerRules.Scan() {
		line := scannerRules.Text()
		rule, err := rule.FromString(line)
		if err != nil {
			return nil, err
		}
		rules[rule.Id] = rule
	}

	return &Repository{
		Facts: facts,
		Rules: rules,
	}, nil
}
