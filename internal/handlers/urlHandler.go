package handlers

import (
	"fmt"
	"itadakimasu-dl/interfaces"
	"itadakimasu-dl/models/animeflv"
	"strings"
)

func GetAnimeByUrl(url string) (interfaces.IAnime, error) {
	if strings.HasPrefix(url, "https://www3.animeflv.net") {
		var name string
		if strings.Contains(url, "/anime/") {
			name = strings.Split(url, "/anime/")[1]
			name = strings.ReplaceAll(name, "-", " ")
		} else if strings.Contains(url, "/ver/") {
			name = strings.Split(url, "/ver/")[1]
			splitted := strings.Split(name, "-")
			name = strings.Join(splitted[:len(splitted)-2], " ")
			url = "https://ww3.animeflv.net/anime/" + strings.ReplaceAll(name, " ", "-")
		} else {
			return nil, fmt.Errorf("AnimeFlv URL format not supported: %s", url)
		}

		return animeflv.NewAnime(1, url), nil
	}
	return nil, fmt.Errorf("URL not supported: %s", url)

}
