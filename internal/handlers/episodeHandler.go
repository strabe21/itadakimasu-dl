package handlers

import (
	"fmt"
	"strconv"
	"strings"
)

func GetEpisodeList(episodeList string) ([]int, error) {
	if episodeList == "0" {
		return nil, nil
	}
	var result []int
	parts := strings.Split(episodeList, ",")
	for _, part := range parts {
		if strings.Contains(part, "-") {
			bounds := strings.Split(part, "-")
			if len(bounds) != 2 {
				return nil, fmt.Errorf("rango inválido: %s", part)
			}

			start, err := strconv.Atoi(bounds[0])
			if err != nil {
				return nil, fmt.Errorf("inicio de rango inválido: %s", bounds[0])
			}

			end, err := strconv.Atoi(bounds[1])
			if err != nil {
				return nil, fmt.Errorf("fin de rango inválido: %s", bounds[1])
			}

			if start > end {
				return nil, fmt.Errorf("el inicio del rango no puede ser mayor que el fin: %s", part)
			}
			if start == 0 {
				return nil, fmt.Errorf("el inicio del rango no puede ser 0")
			}

			for i := start; i <= end; i++ {
				result = append(result, i)
			}
		} else {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("número inválido: %s", part)
			}
			if num == 0 {
				return nil, nil
			}
			result = append(result, num)
		}
	}
	return result, nil
}
