package parser

import (
	"strconv"
	"strings"
)

func ParseId(idStr string) (int, error) {
	idStr = strings.TrimSpace(idStr)
	idStr = strings.TrimLeft(idStr[1:], "0")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return id, nil
}
