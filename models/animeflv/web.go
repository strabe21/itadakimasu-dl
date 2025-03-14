package animeflv

import (
	"fmt"
	"itadakimasu-dl/interfaces"
	"itadakimasu-dl/models"
	"itadakimasu-dl/network"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type animeflv struct {
	models.CrudWeb
}

func (a *animeflv) SearchAnime(query string) ([]interfaces.IAnime, error) {
	formattedQuery := strings.ToLower(strings.ReplaceAll(query, " ", "+"))
	searchURL := a.SearchUrl + url.QueryEscape(formattedQuery)

	doc, err := network.HttpGetDocument(searchURL)
	if err != nil {
		return nil, err
	}
	var animes []interfaces.IAnime
	animeid := 1
	doc.Find("div.Container ul.ListAnimes li article").Each(func(i int, s *goquery.Selection) {
		name := s.Find("h3").Text()

		url, exists := s.Find("a").Attr("href")
		if exists {
			url = a.GetBaseURL() + url
		}

		if name != "" && exists {
			animes = append(animes, NewAnime(animeid, url))
			animeid++
		}
	})
	if len(animes) == 0 {
		return nil, fmt.Errorf("No results found for the search of %s", query)
	}

	return animes, nil
}

func NewAnimeFlv() interfaces.IWeb {
	return &animeflv{models.CrudWeb{
		Name:      "AnimeFlv",
		SearchUrl: "https://www3.animeflv.net/browse?q=",
		BaseUrl:   "https://www3.animeflv.net",
	}}
}
