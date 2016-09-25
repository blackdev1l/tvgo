package tvgo

import (
	"fmt"
	"github.com/apcera/termtables"
	str "github.com/blackdev1l/tvgo/lib/jsonStruct"
	"strconv"
	"strings"
)

func (t *Tvgo) renderSearch(shows []str.Shows) {
	table := termtables.CreateTable()

	fmt.Printf("Found %v series\n", len(shows))

	table.AddHeaders("ID", "Name", "Genre", "Premiered", "Status")
	for _, v := range shows {
		table.AddRow(
			v.ID,
			v.Name,
			strings.Join(v.Genres, ", "),
			v.Premiered,
			v.Status)
	}

	fmt.Println(table.Render())
}

func (t *Tvgo) RenderListShows() {
	t.loadShows()
	table := termtables.CreateTable()
	table.AddHeaders("ID", "Name", "Genre", "Seasons", "Current")
	for _, v := range t.shows {
		lastEpisode := v.Episodes[len(v.Episodes)-1]
		current := findLastSawEpisode(v.Episodes)
		season := strconv.Itoa(current.Season)
		number := strconv.Itoa(current.Number)

		table.AddRow(
			v.ID,
			v.Name,
			strings.Join(v.Genres, ", "),
			lastEpisode.Season,
			season+"x"+number,
		)
	}
	fmt.Println(table.Render())

}

func findLastSawEpisode(episodes []str.Episodes) str.Episodes {
	for _, v := range episodes {
		if v.Seen != true {
			return v
		}
	}

	return episodes[10]
}
