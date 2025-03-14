package animeflv

import (
	"encoding/json"
	"fmt"
	"itadakimasu-dl/interfaces"
	"itadakimasu-dl/models"
	models1 "itadakimasu-dl/models/links"
	"itadakimasu-dl/network"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

type Episode struct {
	models.CrudEpisode
}

func (e *Episode) FetchDownloadLinks() error {
	doc, err := network.HttpGetDocument(e.Url)
	if err != nil {
		return err
	}
	err = e.fetchStreamLinks(doc)
	if err != nil {
		return err
	}
	err = e.fetchDownloadLinks(doc)
	if err != nil {
		return err
	}
	return nil
}
func (e *Episode) fetchStreamLinks(doc *goquery.Document) error {
	var err error = nil
	regex := regexp.MustCompile(`var\s+videos\s*=\s*(\{[\s\S]*?\});`)
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		scriptContent := s.Text()
		match := regex.FindStringSubmatch(scriptContent)
		if len(match) > 1 {
			jsonData := match[1]
			var videos map[string]interface{}
			err = json.Unmarshal([]byte(jsonData), &videos)
			if err != nil {
				return
			}
			for key, videoList := range videos {

				videosSlice, ok := videoList.([]interface{})
				if !ok {
					err = fmt.Errorf("the data type is not a slice for key %s", key)
					continue
				}

				for _, video := range videosSlice {
					videoMap, ok := video.(map[string]interface{})
					if !ok {
						err = fmt.Errorf("the data type is not a map")
						continue
					}
					server, _ := videoMap["title"].(string)
					url, _ := videoMap["url"].(string)
					if url == "" {
						url, _ = videoMap["code"].(string)
					}

					defaultInstance := models1.NewDefaultLink(url, server)
					e.Links = append(e.Links, defaultInstance)
				}
			}
		}
	})
	return err
}
func (e *Episode) fetchDownloadLinks(doc *goquery.Document) error {
	var err error
	doc.Find("table.Dwnl tbody tr").Each(func(i int, row *goquery.Selection) {
		name := row.Find("td").First().Text()
		url, exists := row.Find("td a").Attr("href")
		if !exists {
			err = fmt.Errorf("no se encontró el enlace para la fila: %d", i)
			return
		}
		link := models1.NewDefaultLink(url, name)
		if !e.LinkExist(link) {
			e.Links = append(e.Links, models1.NewDefaultLink(url, name))
		}
	})
	return err
}
func NewEpisode(url string, animeName string, number int) interfaces.IEpisode {
	episode := &Episode{models.CrudEpisode{
		Number:    number,
		Url:       url,
		AnimeName: animeName,
		Links:     make([]interfaces.ILink, 0),
	}}
	episode.FetchDownloadLinks()
	return episode

}
