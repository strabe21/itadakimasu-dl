package handlers

import (
	"itadakimasu-dl/config"
	"itadakimasu-dl/interfaces"
	"itadakimasu-dl/models/animeflv"
)

func GetAnimeWebs() []interfaces.IWeb {
	result := make([]interfaces.IWeb, 0)
	for _, webName := range config.GetConfig.Webs {
		switch webName {
		case "animeflv":
			result = append(result, animeflv.NewAnimeFlv())
		}
	}
	return result
}
