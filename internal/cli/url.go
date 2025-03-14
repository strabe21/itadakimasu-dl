package cli

import (
	"fmt"
	"itadakimasu-dl/internal/handlers"
	"itadakimasu-dl/utils"
	"os"

	"github.com/spf13/cobra"
)

var urlCmd = &cobra.Command{
	Use:   "url [anime_url]",
	Short: "Download an anime using its URL",
	Long:  `Allows to download an anime directly from its URL.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runURL(args[0])
	},
}

func init() {
	urlCmd.Flags().StringVarP(&episodesStr, "episodes", "e", "", "List of episodes to download (example: 1-20 or 1,2). Use 0 for all")
	urlCmd.Flags().StringVarP(&outputDir, "output", "o", "", "Folder where the downloaded episodes will be saved")
	RootCmd.AddCommand(urlCmd)
}

func runURL(animeURL string) {
	utils.PrintAsciiLogo()
	anime, err := handlers.GetAnimeByUrl(animeURL)
	if err != nil {
		fmt.Printf("Error on get anime from url: %v", err)
		os.Exit(1)
	}

	processAnime(anime)
}
