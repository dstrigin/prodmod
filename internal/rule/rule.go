package rule

import (
	"prodmod/internal/fact"
	idpkg "prodmod/pkg/id"
	"strconv"
	"strings"
)

type Rule struct {
	Id          int
	From        []fact.ID
	Result      fact.ID
	Description string
	Weight      float64
}

func FromString(ruleString string) (*Rule, error) {
	tokens := strings.Split(ruleString, ";")
	id, err := idpkg.ParseId(tokens[0])
	if err != nil {
		return nil, err
	}

	from := make([]fact.ID, 0)
	fromRules := strings.Split(tokens[1], ",")
	for _, fromRule := range fromRules {
		ruleId, err := idpkg.ParseId(fromRule)
		if err != nil {
			return nil, err
		}
		from = append(from, ruleId)
	}

	resId, err := idpkg.ParseId(tokens[2])
	if err != nil {
		return nil, err
	}

	description := strings.TrimSpace(tokens[3])

	weight, err := strconv.ParseFloat(tokens[4], 64)
	if err != nil {
		return nil, err
	}

	return &Rule{
		Id:          id,
		From:        from,
		Result:      resId,
		Description: description,
		Weight:      weight,
	}, nil
}
