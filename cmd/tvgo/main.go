package main

import (
	"github.com/blackdev1l/tvgo/lib/config"
	"github.com/blackdev1l/tvgo/lib/tvgo"
	"os"
	"time"

	"github.com/urfave/cli"
)

var (
	seriesName string
	daemonize  bool
	lists      bool
	season     int
	episode    int
	id         int
)

func main() {
	c := config.Configuration{}
	c.CheckConfig()
	tv := tvgo.Tvgo{}
	tv.Conf = c

	app := cli.NewApp()
	app.Name = "Tvgo"
	app.Version = "0.1.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Cristian Achille",
			Email: "blackdev1l@autistici.org",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "search, Ss",
			Usage:       "Lookup the series in the tvmaze database",
			Destination: &seriesName,
		},
		cli.IntFlag{
			Name:        "add, a",
			Usage:       "Add the series in the local database",
			Destination: &id,
		},
		cli.IntFlag{
			Name:        "remove, Rs",
			Usage:       "Remove the series in the local database",
			Destination: &id,
		},
		cli.BoolFlag{
			Name:  "list, ls",
			Usage: "List the series in the local database",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.GlobalIsSet("search") {
			tv.SearchSeries(seriesName)
		} else if c.GlobalIsSet("add") {
			tv.AddSeries(id)
		} else if c.GlobalIsSet("remove") {
			tv.RemoveSeries(id)
		} else if c.GlobalBool("list") {
			tv.RenderListShows()
		}

		return nil
	}

	app.Run(os.Args)

}
