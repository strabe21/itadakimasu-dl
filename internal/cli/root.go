package cli

import (
	"fmt"
	"itadakimasu-dl/config"
	"itadakimasu-dl/interfaces"
	"itadakimasu-dl/internal/handlers"
	"itadakimasu-dl/ui"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	episodesStr string
	outputDir   string
)
var RootCmd = &cobra.Command{
	Use:   "itadakimasu-dl",
	Short: "Itadakimasu DL - Download anime episodes from AnimeFLV (more platforms will be added)",
	Long:  `Itadakimasu DL allows you to search and download anime episodes from AnimeFlv (more platforms will be added in the future) using specific commands.`,
}

func processAnime(anime interfaces.IAnime) {
	fmt.Println("Extracting Episodes and download links...")
	anime.FetchEpisodes()
	if episodesStr != "" && episodesStr != "0" {
		episodeList, err := handlers.GetEpisodeList(episodesStr)
		if err != nil {
			fmt.Printf("error on get episode list: %v\n", err)
			return
		}
		if episodeList != nil {
			anime.SetEpisodesByList(episodeList)
		}
	}
	if outputDir == "" {
		outputDir = config.GetConfig.GetDownloadPath(anime.GetName())
	} else if !strings.HasSuffix(outputDir, "/") {
		outputDir += "/"
	}
	fmt.Print("\033[H\033[2J")
	go ui.GetProgressWriter.Render()
	fmt.Printf("Starting download Anime: %s...\n", anime.GetName())
	anime.Download(outputDir)

	time.Sleep(500 * time.Millisecond)

	bLinks := fmt.Sprint(config.GetConfig.BrokenEpisodes)
	if len(bLinks) > 0 {
		bLinks = strings.Trim(bLinks, "[]")
		if bLinks != "" {
			bLinks = strings.ReplaceAll(bLinks, " ", ",")
			fmt.Printf("Cant download episodes %s, the links are broken\n", bLinks)
		}
	}
	fmt.Println("All downloads finished, itadakimasu!")
	ui.GetProgressWriter.Stop()
	return
}
