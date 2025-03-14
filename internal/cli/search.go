package cli

import (
	"fmt"
	"itadakimasu-dl/internal/handlers"
	"itadakimasu-dl/utils"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [anime_name]",
	Short: "Search for an anime by name",
	Long:  `Allows you to search for an anime by name and download the selected episodes.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runSearch(args[0])
	},
}

func init() {
	searchCmd.Flags().StringVarP(&episodesStr, "episodes", "e", "", "List of episodes to download (example: 1-20 or 1,2). Use 0 for all")
	searchCmd.Flags().StringVarP(&outputDir, "output", "o", "", "Folder where the downloaded episodes will be saved")
	RootCmd.AddCommand(searchCmd)
}

func runSearch(animeName string) {
	utils.PrintAsciiLogo()
	anime := handlers.Search(animeName)
	if anime == nil {
		fmt.Println("Anime not found.")
		return
	}
	processAnime(anime)
}
