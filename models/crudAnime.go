package models

import (
	"errors"
	"fmt"
	"itadakimasu-dl/interfaces"
	"strings"
	"sync"
)

type CrudAnime struct {
	Id       int
	Url      string
	Name     string
	Episodes map[int]interfaces.IEpisode
}

func (a *CrudAnime) Download(animePath string) {
	var wg sync.WaitGroup
	for _, episode := range a.Episodes {
		wg.Add(1)
		go episode.InitDownload(animePath, &wg)
	}
	wg.Wait()
}

func (a *CrudAnime) GetEpisodes() map[int]interfaces.IEpisode {
	return a.Episodes
}

func (a *CrudAnime) GetId() int {
	return a.Id
}

func (a *CrudAnime) GetName() string {
	return a.Name
}

func (a *CrudAnime) GetUrl() string {
	return a.Url
}

func (a *CrudAnime) GetEpisodeByNumber(number int) (interfaces.IEpisode, error) {
	if episode, exists := a.Episodes[number]; exists {
		return episode, nil
	}
	return nil, fmt.Errorf("episode with number %d does not exist", number)
}

func (a *CrudAnime) SetEpisodesByList(list []int) error {
	result := make(map[int]interfaces.IEpisode)
	var errorsList []string
	for _, i := range list {
		epi, err := a.GetEpisodeByNumber(i)
		if err != nil {
			errorsList = append(errorsList, err.Error())
			continue
		}
		result[i] = epi
	}
	if len(errorsList) > 0 {
		return errors.New(strings.Join(errorsList, ", "))
	}
	a.Episodes = result
	return nil
}
