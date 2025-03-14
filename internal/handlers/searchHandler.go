package handlers

import (
	"itadakimasu-dl/interfaces"
	"itadakimasu-dl/ui"
	"itadakimasu-dl/utils"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func Search(query string) interfaces.IAnime {
	results := make(map[int]interfaces.IAnime)
	curId := 1
	for _, web := range GetAnimeWebs() {
		webResult, err := web.SearchAnime(query)
		if err != nil {
			continue
		}
		for _, wr := range webResult {
			results[curId] = wr
			curId++
		}
	}

	tHeader := table.Row{"ID", "Name", "Source"}
	tConfig := []table.ColumnConfig{
		{Name: "ID", WidthMin: 10, Align: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Name: "Name", WidthMin: 50, Align: text.AlignCenter, AlignHeader: text.AlignCenter},
		{Name: "Source", WidthMin: 15, Align: text.AlignCenter, AlignHeader: text.AlignCenter},
	}

	tRows := []table.Row{}
	for _, anime := range results {
		tRows = append(tRows, table.Row{anime.GetId(), anime.GetName(), anime.GetSource()})
	}
	ui.PrintTable("Search results from: "+query, tHeader, tRows, nil, tConfig)

	selected := utils.GetIntFromTerm("Select an anime to download (0 to exit)")
	if selected == 0 {
		os.Exit(0)

	}
	return results[selected]
}
