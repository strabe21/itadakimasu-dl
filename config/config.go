package config

import (
	"encoding/json"
	"fmt"
	"itadakimasu-dl/utils"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/sync/semaphore"
)

type Config struct {
	DownloadPath   string             `json:"downloadPath"`
	AnimePath      string             `json:"animePath"`
	EpisodeFile    string             `json:"episodeFile"`
	Webs           []string           `json:"animeWebs"`
	Links          []map[string]*Link `json:"links"`
	BrokenEpisodes []int
}

var GetConfig *Config = getConfig()

func getConfigPath() string {
	exePath, err := os.Executable()
	if err != nil {
		return "config.json"
	}

	configPath := filepath.Join(filepath.Dir(exePath), "config.json")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return "config.json"
	}

	return configPath
}
func getConfig() *Config {
	file, err := os.Open(getConfigPath())
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}
	for _, linkMap := range config.Links {
		for _, linkConfig := range linkMap {
			linkConfig.Semaphore = semaphore.NewWeighted(int64(linkConfig.MaxConcurrentDownloads))
		}
	}
	return &config
}

func (c *Config) GetDownloadPath(AnimeName string) string {
	dDir := c.DownloadPath
	dDir = strings.ReplaceAll(dDir, "\\", "/")
	if !strings.HasSuffix(dDir, "/") {
		dDir += "/"
	}
	splitedDpath := strings.Split(c.AnimePath, "/")
	re := regexp.MustCompile(`\[([^\[\]]+)\]`)
	for _, dir := range splitedDpath {

		matches := re.FindAllStringSubmatch(dir, -1)

		for _, match := range matches {
			content := match[1]

			switch content {
			case "NAME":
				dDir += strings.ReplaceAll(dir, match[0], AnimeName)
			case "ASK":
				res, _ := utils.AskOnTerm("", dDir)
				return res
			default:
				fmt.Printf("What the hell is: %s\n", content)
			}
		}
		if !strings.HasSuffix(dDir, "/") {
			dDir += "/"
		}
	}
	return dDir
}

func (c *Config) GetLinkConfig(name string) *Link {
	for _, linksMap := range c.Links {
		for linkName, linkConf := range linksMap {
			if linkName == name {
				return linkConf
			}
		}
	}
	return nil
}
