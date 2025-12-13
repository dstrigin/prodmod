package repository

import (
	"bufio"
	"os"
	"prodmod/internal/fact"
	"prodmod/internal/rule"
)

type Repository struct {
	Facts []fact.Fact
	Rules []rule.Rule
}

func NewRepository(factsFileName, rulesFileName string) (*Repository, error) {
	factsFile, err := os.Open(factsFileName)
	if err != nil {
		return nil, err
	}
	defer factsFile.Close()

	facts := make([]fact.Fact, 0)
	scannerFacts := bufio.NewScanner(factsFile)
	for scannerFacts.Scan() {
		line := scannerFacts.Text()
		fact, err := fact.FromString(line)
		if err != nil {
			return nil, err
		}
		facts = append(facts, *fact)
	}

	rulesFile, err := os.Open(rulesFileName)
	if err != nil {
		return nil, err
	}
	defer rulesFile.Close()

	rules := make([]rule.Rule, 0)
	scannerRules := bufio.NewScanner(rulesFile)
	for scannerRules.Scan() {
		line := scannerRules.Text()
		rule, err := rule.FromString(line)
		if err != nil {
			return nil, err
		}
		rules = append(rules, *rule)
	}

	return &Repository{
		Facts: facts,
		Rules: rules,
	}, nil
}
