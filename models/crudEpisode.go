package models

import (
	"fmt"
	"itadakimasu-dl/config"
	"itadakimasu-dl/interfaces"
	"itadakimasu-dl/utils"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type CrudEpisode struct {
	Number    int
	Url       string
	AnimeName string
	Links     []interfaces.ILink
}

func (e *CrudEpisode) DownloadLink(link interfaces.ILink, outputFile string) (bool, error) {
	linkConfig := link.GetConfig()

	sem := linkConfig.Semaphore
	acquired := sem.TryAcquire(1)
	if !acquired {
		return false, nil
	}
	defer sem.Release(1)

	err := link.Download(outputFile)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (e *CrudEpisode) InitDownload(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	err := os.MkdirAll(path, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	busyLinks := make([]interfaces.ILink, 0)
	outputFile := path + e.GetOutputFile()
	for prio := 1; prio < 3; prio++ {
		for _, link := range e.Links {

			if link.GetConfig().Priority == prio {
				success, err := e.DownloadLink(link, outputFile)
				if err != nil {
					continue
				}
				if !success {
					busyLinks = append(busyLinks, link)
				} else {
					return
				}
			}

		}

	}
	for len(busyLinks) > 0 {
		for i := len(busyLinks) - 1; i >= 0; i-- {
			link := busyLinks[i]
			success, err := e.DownloadLink(link, outputFile)
			if err != nil {
				busyLinks = utils.RemoveFromSlice(busyLinks, i)
				continue
			}
			if success {
				return
			}
		}
		time.Sleep(5 * time.Second)
	}
	fmt.Println("failed")
	config.GetConfig.BrokenEpisodes = append(config.GetConfig.BrokenEpisodes, e.Number)
}
func (e *CrudEpisode) GetOutputFile() string {
	var result string = config.GetConfig.EpisodeFile
	re := regexp.MustCompile(`\[([^\[\]]+)\]`)
	matches := re.FindAllStringSubmatch(result, -1)
	for _, match := range matches {
		content := match[1]

		switch content {
		case "NAME":
			result = strings.ReplaceAll(result, match[0], e.AnimeName)
		case "ASK":
			res, _ := utils.AskOnTerm("Episode name", strings.ReplaceAll(config.GetConfig.EpisodeFile, match[0], ""))
			return res
		case "NUMBER":
			result = strings.ReplaceAll(result, match[0], strconv.Itoa(e.Number))
		default:
			fmt.Printf("What the hell is: %s\n", content)
			result = strconv.Itoa(e.Number)
		}
	}
	result += ".mp4"
	return result
}
func (e *CrudEpisode) LinkExist(link interfaces.ILink) bool {
	for _, l := range e.Links {
		if l.GetServerName() == link.GetServerName() && l.GetUrl() == link.GetUrl() {
			return true
		}
	}
	return false
}
func (e *CrudEpisode) GetAnimeName() string {
	return e.AnimeName
}
func (e *CrudEpisode) GetLinks() []interfaces.ILink {
	return e.Links
}
func (e *CrudEpisode) GetNumber() int {
	return e.Number
}

func (e *CrudEpisode) GetUrl() string {
	return e.Url
}
