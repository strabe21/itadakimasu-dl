package animeflv

import (
	"fmt"
	"itadakimasu-dl/interfaces"
	"itadakimasu-dl/models"
	"itadakimasu-dl/network"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Anime struct {
	models.CrudAnime
}

func (a *Anime) GetSource() string {
	return "AnimeFlv"
}

func (a *Anime) FetchEpisodes() error {
	doc, err := network.HttpGetDocument(a.Url)
	if err != nil {
		return err
	}
	var scriptText string
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		if s.Text() != "" && regexp.MustCompile(`episodes`).MatchString(s.Text()) {
			scriptText = s.Text()
		}
	})
	if scriptText == "" {
		return fmt.Errorf("script with episodes information of anime %s not found", a.Name)
	}

	episodeRegex := regexp.MustCompile(`\[(\d+),(\d+)\]`)
	matches := episodeRegex.FindAllStringSubmatch(scriptText, -1)

	baseUrl := strings.ReplaceAll(a.Url, "/anime/", "/ver/")
	for _, match := range matches {
		episodeNumber, err := strconv.Atoi(match[1])
		if err != nil {
			return fmt.Errorf("error converting episode number: %w", err)
		}
		episodeURL := fmt.Sprintf("%s-%d", baseUrl, episodeNumber)
		a.Episodes[episodeNumber] = NewEpisode(episodeURL, a.Name, episodeNumber)
	}
	return nil
}
func NewAnime(id int, url string) interfaces.IAnime {
	name := strings.ReplaceAll(strings.Split(url, "/anime/")[1], "-", " ")
	anime := &Anime{models.CrudAnime{
		Id:       id,
		Url:      url,
		Name:     name,
		Episodes: make(map[int]interfaces.IEpisode),
	}}
	return anime
}
