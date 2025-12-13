package id

import (
	"bufio"
	"os"
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

func ReadIds(fileName string) ([]int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	ids := make([]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		id, err := ParseId(line)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}
