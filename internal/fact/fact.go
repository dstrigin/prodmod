package fact

import (
	"prodmod/pkg/id"
	"strconv"
	"strings"
)

type ID = int

type Fact struct {
	Id          ID
	Description string
	Weight      float64
}

func FromString(factString string) (*Fact, error) {
	tokens := strings.Split(factString, ";")
	id, err := id.ParseId(tokens[0])
	if err != nil {
		return nil, err
	}

	description := tokens[1]
	weight, err := strconv.ParseFloat(tokens[2], 64)
	if err != nil {
		return nil, err
	}

	return &Fact{
		Id:          id,
		Description: description,
		Weight:      weight,
	}, nil
}
